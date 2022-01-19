package client

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type idsFlag []string

func (list idsFlag) String() string {
	return strings.Join(list, ",")
}

func (list *idsFlag) Set(v string) error {
	*list = append(*list, v)
	return nil
}

type BackendHTTPClient interface {
	Create(title, message string, duration time.Duration)([]byte, error)
	Edit(id, title, message string, duration time.Duration)([]byte, error)
	Fetch(ids []string)([]byte, error)
	Delete(ids []string) error
	Healthy(host string) bool
}

func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	s := Switch{
		client: httpClient,
		backendAPIURL: uri,
	}
	s.commands = map[string]func() func(string) error{
		"create": s.create,
		"edit":   s.edit,
		"fetch":  s.fetch,
		"delete": s.delete,
		"health": s.health,
	}
	return s
}
type Switch struct {
	client        BackendHTTPClient
	backendAPIURL string
	commands      map[string]func() func(string) error
}
func (s Switch) Switch() error{
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]
	if !ok{
		return fmt.Errorf("invalid command - '%s'", cmdName)
	}
	return cmd()(cmdName) 
}
func (s Switch) Help() {
	var help string
	for name := range s.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("Usage: %s: <command> [<args>]\n%s", os.Args[0], help)
}

func (s Switch) create() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m , d := s.reminderFlags(createCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}
		if err := s.parseCmd(createCmd); err != nil {
			return err
		}
		res, err := s.client.Create(*t, *m,*d)
		if err != nil {
			return wrapError("Could not create Reminder", err)
		}
		fmt.Printf("Reminder created successfully: \n%s", string(res))
		return nil
	}
}

func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		editCmd.Var(&ids, "id", "The Id of edit reminder")
		t, m , d := s.reminderFlags(editCmd)

		if err := s.checkArgs(2); err != nil {
			return err
		}
		if err := s.parseCmd(editCmd); err != nil {
			return err
		}

		lastID := ids[len(ids) - 1]
		res, err := s.client.Edit(lastID, *t, *m,*d)
		if err != nil {
			return wrapError("Could not edit Reminder", err)
		}
		fmt.Printf("Reminder edited successfully: \n%s", string(res))
		return nil
	}
}
func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fetchCmd.Var(&ids, "id", "The Id of fetch reminder")
		
		if err := s.checkArgs(1); err != nil {
			return err
		}
		if err := s.parseCmd(fetchCmd); err != nil {
			return err
		}

		
		res, err := s.client.Fetch(ids)
		if err != nil {
			return wrapError("Could not fetch Reminder", err)
		}
		fmt.Printf("Reminder fetched successfully: \n%s", string(res))
		return nil
	}
}
func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		deleteCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		deleteCmd.Var(&ids, "id", "List of reminder IDs to delete")
		
		if err := s.checkArgs(1); err != nil {
			return err
		}
		if err := s.parseCmd(deleteCmd); err != nil {
			return err
		}

		
		err := s.client.Delete(ids)
		if err != nil {
			return wrapError("Could not delete Reminder", err)
		}
		fmt.Printf("Reminder deleted successfully: \n%v\n", ids)
		return nil
	}
}
func (s Switch) health() func(string) error {
	return func(cmd string) error {
		var host string
		healthCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		healthCmd.StringVar(&host, "host", s.backendAPIURL, "ping for health")
		if err := s.parseCmd(healthCmd); err != nil {
			return err
		}
		if !s.client.Healthy(host){
			fmt.Printf("host %s is running\n", host)
		}else {
			fmt.Printf("host %s is down\n",0ost)
		}
		
		return nil
	}
}

func (s Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, m, d := "", "",time.Duration(0)
	f.StringVar(&t, "title", "", "Title")
	f.StringVar(&t, "t", "", "Title")
	f.StringVar(&m, "message", "", "Message")
	f.StringVar(&m, "m", "", "Message")
	f.StringVar(&d, "duration",0,"Duration")
	f.StringVar(&d, "d", 0, "Duration")
	return &t, &m, &d
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError("Couldnot parse successfully '"+cmd.Name()+"' flags", err)
	}
	return nil
}

func (s Switch) checkArgs(minArgs int) error{
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}
	if len(os.Args) - 2 < minArgs {
		fmt.Printf("Incorrect use of %s\n%s %s --help\n", os.Args[1],os.Args[0],os.Args[1])
		return fmt.Errorf("%s expects atleast %d arg(s), %d provided", 
			os.Args[1], minArgs, len(os.Args)-2,
		)
	}
	return nil
}