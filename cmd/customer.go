/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"

	"github.com/google/uuid"

	"github.com/TemurMannonov/buy_event/models"
	"github.com/spf13/cobra"
)

// customerCmd represents the customer command
var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Command for getting, adding, updating and deleting customer.",
	Long:  "Command for getting, adding, updating and deleting customer.",
}

var addCustomerCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new customer",
	RunE: func(cmd *cobra.Command, args []string) error {
		phoneNumber, err := cmd.Flags().GetString("phone")
		if err != nil {
			return err
		}

		if phoneNumber == "" {
			return errors.New("Phone number is required")
		}

		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}

		if email == "" {
			return errors.New("Email is required")
		}

		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		err = strg.Customer().Create(&models.Customer{
			ID:          id.String(),
			Email:       email,
			PhoneNumber: phoneNumber,
		})

		if err != nil {
			return err
		}

		cmd.Println("Customer successfully created!")

		return nil
	},
}

var getCustomerCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a customer info",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		if id == "" {
			customers, err := strg.Customer().GetAll()
			if err != nil {
				return err
			}

			cmd.Println(`Customers List`)
			for _, customer := range customers {
				cmd.Println(`-----------------------------------`)
				cmd.Println("ID: ", customer.ID)
				cmd.Println("Email: ", customer.Email)
				cmd.Println("Phone: ", customer.PhoneNumber)
			}
		} else {
			_, err := uuid.Parse(id)
			if err != nil {
				return err
			}

			customer, err := strg.Customer().Get(id)
			if err != nil {
				return err
			}

			cmd.Println(`Customer Information`)
			cmd.Println(`-----------------------------------`)
			cmd.Println("ID: ", customer.ID)
			cmd.Println("Email: ", customer.Email)
			cmd.Println("Phone: ", customer.PhoneNumber)
		}

		return nil
	},
}

var updateCustomerCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a customer info",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		_, err = uuid.Parse(id)
		if err != nil {
			return err
		}

		phoneNumber, err := cmd.Flags().GetString("phone")
		if err != nil {
			return err
		}

		if phoneNumber == "" {
			return errors.New("Phone number is required")
		}

		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}

		if email == "" {
			return errors.New("Email is required")
		}

		err = strg.Customer().Update(&models.Customer{
			ID:          id,
			Email:       email,
			PhoneNumber: phoneNumber,
		})

		if err != nil {
			return err
		}

		cmd.Println("Customer successfully updated!")
		return nil
	},
}

var deleteCustomerCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a customer",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		_, err = uuid.Parse(id)
		if err != nil {
			return err
		}

		err = strg.Customer().Delete(id)
		if err != nil {
			return err
		}

		cmd.Println("Customer successfully deleted")
		return nil
	},
}

func init() {
	addCustomerCmd.Flags().StringP("phone", "p", "", "customer phone number")
	addCustomerCmd.Flags().StringP("email", "e", "", "customer email")
	customerCmd.AddCommand(addCustomerCmd)

	getCustomerCmd.Flags().StringP("id", "i", "", "customer id")
	customerCmd.AddCommand(getCustomerCmd)

	updateCustomerCmd.Flags().StringP("id", "i", "", "customer id")
	updateCustomerCmd.Flags().StringP("email", "e", "", "customer email")
	updateCustomerCmd.Flags().StringP("phone", "p", "", "customer phone number")
	customerCmd.AddCommand(updateCustomerCmd)

	deleteCustomerCmd.Flags().StringP("id", "i", "", "customer id")
	customerCmd.AddCommand(deleteCustomerCmd)

	rootCmd.AddCommand(customerCmd)
}
