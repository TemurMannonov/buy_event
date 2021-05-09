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
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Customer command")
	},
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

			for _, customer := range customers {
				cmd.Println(`----------**********----------`)
				cmd.Println("ID: ", customer.ID)
				cmd.Println("Email: ", customer.Email)
				cmd.Println("Phone: ", customer.PhoneNumber)
			}
		} else {
			customerID, err := uuid.Parse(id)
			if err != nil {
				return err
			}

			customer, err := strg.Customer().Get(customerID.String())
			if err != nil {
				return err
			}

			cmd.Println("ID: ", customer.ID)
			cmd.Println("Email: ", customer.Email)
			cmd.Println("Phone: ", customer.PhoneNumber)
		}

		return nil
	},
}

var updateCustomerCmd = &cobra.Command{
	Use:   "add",
	Short: "Update a customer info",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
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

		customerID, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		err = strg.Customer().Delete(customerID.String())
		if err != nil {
			return err
		}

		cmd.Println("Customer successfully deleted")
		return nil
	},
}

func init() {
	addCustomerCmd.Flags().StringP("email", "e", "", "Enter customer email")
	addCustomerCmd.Flags().StringP("phone", "p", "", "Enter customer phone")

	getCustomerCmd.Flags().StringP("id", "i", "", "Enter customer ID")

	updateCustomerCmd.Flags().StringP("email", "e", "", "Enter customer email")
	updateCustomerCmd.Flags().StringP("phone", "p", "", "Enter customer phone")

	deleteCustomerCmd.Flags().StringP("id", "i", "", "Enter customer id")

	customerCmd.AddCommand(addCustomerCmd)
	customerCmd.AddCommand(getCustomerCmd)
	customerCmd.AddCommand(updateCustomerCmd)
	customerCmd.AddCommand(deleteCustomerCmd)

	rootCmd.AddCommand(customerCmd)
}
