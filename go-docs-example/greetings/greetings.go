package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Function to greet a single person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name");
	}
	// var message string, message = || shortcut -> := (interpret and assign)
	message := fmt.Sprintf(randomFormat(), name);
	return message, nil;
}

// Functiont to greet multiple people
func Hellos(names []string) (map[string] string, error) {
	messages := make(map[string] string);

	for _, name := range names {
		message, err := Hello(name);
		if err != nil {
			return nil, err;
		}

		messages[name] = message;
	}

	return messages, nil;
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	};

	return formats[rand.Intn(len(formats))];
}