package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/olimci/prompter"
)

func ccnValidator(s string) error {
	// Credit Card Number should a string less than 20 digits
	// It should include 16 integers and 3 spaces
	if len(s) > 16+3 {
		return fmt.Errorf("CCN is too long")
	}

	if len(s) == 0 || len(s)%5 != 0 && (s[len(s)-1] < '0' || s[len(s)-1] > '9') {
		return fmt.Errorf("CCN is invalid")
	}

	// The last digit should be a number unless it is a multiple of 4 in which
	// case it should be a space
	if len(s)%5 == 0 && s[len(s)-1] != ' ' {
		return fmt.Errorf("CCN must separate groups with spaces")
	}

	// The remaining digits should be integers
	c := strings.ReplaceAll(s, " ", "")
	_, err := strconv.ParseInt(c, 10, 64)

	return err
}

func expValidator(s string) error {
	// The 3 character should be a slash (/)
	// The rest should be numbers
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("EXP is invalid")
	}

	// There should be only one slash and it should be in the 2nd index (3rd character)
	if len(s) >= 3 && (strings.Index(s, "/") != 2 || strings.LastIndex(s, "/") != 2) {
		return fmt.Errorf("EXP is invalid")
	}

	return nil
}

func cvvValidator(s string) error {
	// The CVV should be a number of 3 digits
	// Since the input will already ensure that the CVV is a string of length 3,
	// All we need to do is check that it is a number
	_, err := strconv.ParseInt(s, 10, 64)
	return err
}

func main() {
	ctx := context.Background()

	err := prompter.Start(func(p *prompter.Prompter) error {
		p.Log("Credit Card Form")

		cardType, err := p.AwaitSelectDefault("Card type:", []string{"Visa", "Mastercard", "Amex"}, "Amex")
		if err != nil {
			return err
		}
		p.Logf("card type: %s", cardType)

		ccnPro, err := p.Input(
			prompter.WithInputPrompt("CCN: "),
			prompter.WithInputValidate(ccnValidator),
		)
		if err != nil {
			return err
		}

		expPro, err := p.Input(
			prompter.WithInputPrompt("EXP: "),
			prompter.WithInputValidate(expValidator),
		)
		if err != nil {
			return err
		}

		cvvPro, err := p.Input(
			prompter.WithInputPrompt("CVV: "),
			prompter.WithInputValidate(cvvValidator),
		)
		if err != nil {
			return err
		}

		ccn, err := ccnPro.Await(p.Ctx)
		if err != nil {
			return err
		}

		exp, err := expPro.Await(p.Ctx)
		if err != nil {
			return err
		}

		cvv, err := cvvPro.Await(p.Ctx)
		if err != nil {
			return err
		}

		p.Logf("CCN: %s", ccn)
		p.Logf("EXP: %s", exp)
		p.Logf("CVV: %s", cvv)

		return nil
	},
		prompter.WithContext(ctx), // whatever your Option is called; adjust if different
	)

	if err != nil {
		panic(err)
	}
}
