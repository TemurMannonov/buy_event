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

	"github.com/TemurMannonov/buy_event/config"
	"github.com/TemurMannonov/buy_event/helpers"
	"github.com/TemurMannonov/buy_event/models"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Command for getting, adding, updating and deleting order.",
	Long:  "Command for getting, adding, updating and deleting order.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Order command")
	},
}

var addOrderCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new order",
	RunE: func(cmd *cobra.Command, args []string) error {
		customerID, err := cmd.Flags().GetString("customer-id")
		if err != nil {
			return err
		}

		if customerID == "" {
			return errors.New("Customer id is required")
		}

		_, err = uuid.Parse(customerID)
		if err != nil {
			return err
		}

		totalPrice, err := cmd.Flags().GetFloat64("total-price")
		if err != nil {
			return err
		}

		if totalPrice < 0 {
			return errors.New("Total price should be greater than 0")
		}

		products, err := cmd.Flags().GetString("products")
		if err != nil {
			return err
		}

		if products == "" {
			return errors.New("Products are required")
		}

		notificationType, err := cmd.Flags().GetString("notification-type")
		if err != nil {
			return err
		}

		if !validateNotificationType(notificationType) {
			return errors.New("Notification type is incorrect")
		}

		err = strg.Order().Create(&models.Order{
			Products:   products,
			CustomerID: customerID,
			TotalPrice: totalPrice,
		})

		if err != nil {
			return err
		}

		cmd.Println("Order successfully created!")

		customer, err := strg.Customer().Get(customerID)

		go sendNotification(notificationType, customer)

		return nil
	},
}

var getOrderCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an order info",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")

		if err != nil {
			return err
		}

		if id == "" {
			orders, err := strg.Order().GetAll()

			if err != nil {
				return err
			}

			for _, order := range orders {
				cmd.Println(`----------**********----------`)
				cmd.Println("ID: ", order.ID)
				cmd.Println("Customer ID: ", order.CustomerID)
				cmd.Println("Products: ", order.Products)
				cmd.Println("Total price: ", order.TotalPrice)
			}
		} else {
			_, err := uuid.Parse(id)
			if err != nil {
				return err
			}

			order, err := strg.Order().Get(id)
			if err != nil {
				return err
			}

			cmd.Println("ID: ", order.ID)
			cmd.Println("Customer ID: ", order.CustomerID)
			cmd.Println("Products: ", order.Products)
			cmd.Println("Total price: ", order.TotalPrice)
		}

		return nil
	},
}

var updateOrderCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an order info",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		_, err = uuid.Parse(id)
		if err != nil {
			return err
		}

		customerID, err := cmd.Flags().GetString("customer-id")
		if err != nil {
			return err
		}

		if customerID == "" {
			return errors.New("Customer id is required")
		}

		_, err = uuid.Parse(customerID)
		if err != nil {
			return err
		}

		totalPrice, err := cmd.Flags().GetFloat64("total-price")
		if err != nil {
			return err
		}

		if totalPrice < 0 {
			return errors.New("Total price should be greater than 0")
		}

		products, err := cmd.Flags().GetString("products")
		if err != nil {
			return err
		}

		if products == "" {
			return errors.New("Products are required")
		}

		err = strg.Order().Update(&models.Order{
			ID:         id,
			Products:   products,
			TotalPrice: totalPrice,
		})

		if err != nil {
			return err
		}

		cmd.Println("Order successfully updated!")

		return nil
	},
}

var deleteOrderCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an order",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		_, err = uuid.Parse(id)
		if err != nil {
			return err
		}

		err = strg.Order().Delete(id)
		if err != nil {
			return err
		}

		cmd.Println("Order successfully deleted!")
		return nil
	},
}

func init() {
	addOrderCmd.Flags().StringP("customer-id", "c", "", "Enter customer id")
	addOrderCmd.Flags().Float64P("total-price", "tp", 0, "Enter total price")
	addOrderCmd.Flags().StringP("products", "p", "", "Enter products")
	addOrderCmd.Flags().StringP("notification-type", "nt", "email", "Notification type")

	getOrderCmd.Flags().StringP("id", "i", "", "Enter order id")

	updateOrderCmd.Flags().StringP("id", "i", "", "Enter order id")
	updateOrderCmd.Flags().Float64P("total-price", "tp", 0, "Enter total price")
	updateOrderCmd.Flags().StringP("products", "p", "", "Enter products")

	deleteOrderCmd.Flags().StringP("id", "i", "", "Enter order id")

	orderCmd.AddCommand(addOrderCmd)
	orderCmd.AddCommand(getOrderCmd)
	orderCmd.AddCommand(getOrderCmd)
	orderCmd.AddCommand(deleteOrderCmd)

	rootCmd.AddCommand(orderCmd)
}

func validateNotificationType(notificationType string) bool {
	for _, val := range config.NotificationTypes {
		if val == notificationType {
			return true
		}
	}
	return false
}

func sendNotification(notificationType string, customer *models.Customer) {
	if notificationType == config.NotificationTypeEmail {
		helpers.SendEmail(customer.Email, "Ok")
	} else if notificationType == config.NotificationTypeSms {
		helpers.SendEmail(customer.PhoneNumber, "Ok")
	}
}
