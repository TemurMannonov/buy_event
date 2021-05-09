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
	"fmt"

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

		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		err = strg.Order().Create(&models.Order{
			ID:         id.String(),
			Products:   products,
			CustomerID: customerID,
			TotalPrice: totalPrice,
		})

		if err != nil {
			return err
		}

		cmd.Println("Order successfully created!")

		customer, err := strg.Customer().Get(customerID)
		if err != nil {
			return err
		}

		text := fmt.Sprintf("Thanks for your oder:"+
			"\nProducts: %s"+
			"\nTotal price: %f",
			products, totalPrice,
		)

		if notificationType == config.NotificationTypeEmail {
			err := helpers.SendEmail(customer.Email, text)
			if err != nil {
				createLog("Error while sending sms: " + err.Error())
				return err
			}

		} else if notificationType == config.NotificationTypeSms {
			err := helpers.SendSMS(customer.PhoneNumber, text)
			if err != nil {
				createLog("Error while sending sms: " + err.Error())
				return err
			}
		}

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

			cmd.Println(`Orders List`)
			for _, order := range orders {
				cmd.Println(`-----------------------------------`)
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

			cmd.Println(`Orders Information`)
			cmd.Println(`-----------------------------------`)
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
			CustomerID: customerID,
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
	addOrderCmd.Flags().StringP("customer-id", "c", "", "customer id")
	addOrderCmd.Flags().Float64P("total-price", "t", 0, "total price")
	addOrderCmd.Flags().StringP("products", "p", "", "products")
	addOrderCmd.Flags().StringP("notification-type", "n", "sms", "notification type")
	orderCmd.AddCommand(addOrderCmd)

	getOrderCmd.Flags().StringP("id", "i", "", "order id")
	orderCmd.AddCommand(getOrderCmd)

	updateOrderCmd.Flags().StringP("id", "i", "", "order id")
	updateOrderCmd.Flags().StringP("customer-id", "c", "", "customer id")
	updateOrderCmd.Flags().Float64P("total-price", "t", 0, "total price")
	updateOrderCmd.Flags().StringP("products", "p", "", "products")
	orderCmd.AddCommand(updateOrderCmd)

	deleteOrderCmd.Flags().StringP("id", "i", "", "order id")
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

func createLog(message string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	err = strg.Log().Create(&models.Log{
		ID:      id.String(),
		Message: message,
	})

	if err != nil {
		return err
	}

	return nil
}
