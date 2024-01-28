/*

@Getter
@Setter
@AllArgsConstructor
@Builder
@NoArgsConstructor
@ToString
@EqualsAndHashCode(of = "lineNumber")
public class LineItem {

  private Item item;
  private String orderedQuantity;
  private RewardsOrderStatus status;
  private Boolean lockedForUpdate;

  @NotNull
  @Min(1)
  private Integer lineNumber;

  private Integer subLineNumber;
  private LinePriceInfo linePriceInfo;
  private List<Note> notes;
  private List<HoldHistory> holdHistory;
  private OrderLineAdditionalInformation additionalInfo;
  private boolean ignoreForPayment;

  public boolean checkIfLineOnHold() {
    return getNonNullList(this.getHoldHistory()).stream().anyMatch(HoldHistory::checkIfHoldCreated);
  }
}


*/



package main

import (
	"time"
)

// RewardsOrderStatus is an enum-like type
type RewardsOrderStatus string

// Note represents a note
type Note struct {
	// Define fields as needed
}

// HoldHistory represents hold history
type HoldHistory struct {
	// Define fields as needed
}

// LinePriceInfo represents line price information
type LinePriceInfo struct {
	// Define fields as needed
}

// OrderLineAdditionalInformation represents additional information for an order line
type OrderLineAdditionalInformation struct {
	// Define fields as needed
}

// LineItem represents a line item in an order
type LineItem struct {
	Item               Item
	OrderedQuantity    string
	Status             RewardsOrderStatus
	LockedForUpdate    *bool
	LineNumber         int `validate:"min=1,nonzero"`
	SubLineNumber      *int
	LinePriceInfo      LinePriceInfo
	Notes              []Note
	HoldHistory        []HoldHistory
	AdditionalInfo     OrderLineAdditionalInformation
	IgnoreForPayment   bool
}

// Item represents an item
type Item struct {
	// Define fields as needed
}

// validate is a placeholder function for validation logic
func (li *LineItem) validate() error {
	// Add validation logic here
	return nil
}

// CheckIfLineOnHold checks if the line item is on hold
func (li *LineItem) CheckIfLineOnHold() bool {
	return li.checkIfHoldCreated()
}

// checkIfHoldCreated checks if any hold history indicates a hold is created
func (li *LineItem) checkIfHoldCreated() bool {
	for _, hold := range li.HoldHistory {
		if hold.checkIfHoldCreated() {
			return true
		}
	}
	return false
}

// checkIfHoldCreated is a placeholder method in HoldHistory
func (hh *HoldHistory) checkIfHoldCreated() bool {
	// Add logic to check if hold is created
	return false
}

// getNonNullList is a placeholder function for handling non-null lists
func (li *LineItem) getNonNullList(list []HoldHistory) []HoldHistory {
	// Add logic to handle non-null lists
	return list
}

func main() {
	// Example usage
	lineItem := LineItem{
		LineNumber: 1,
		// Initialize other fields as needed
	}

	err := lineItem.validate()
	if err != nil {
		// Handle validation error
	}
}



package main

import "fmt"

// RewardsOrderStatus is an enum-like type
type RewardsOrderStatus string

// Note represents a note
type Note struct {
	// Define fields as needed
}

// HoldHistory represents hold history
type HoldHistory struct {
	// Define fields as needed
}

// LinePriceInfo represents line price information
type LinePriceInfo struct {
	// Define fields as needed
}

// OrderLineAdditionalInformation represents additional information for an order line
type OrderLineAdditionalInformation struct {
	// Define fields as needed
}

// LineItem represents a line item in an order
type LineItem struct {
	Item               Item
	OrderedQuantity    string
	Status             RewardsOrderStatus
	LockedForUpdate    *bool
	LineNumber         int `validate:"min=1,nonzero"`
	SubLineNumber      *int
	LinePriceInfo      LinePriceInfo
	Notes              []Note
	HoldHistory        []HoldHistory
	AdditionalInfo     OrderLineAdditionalInformation
	IgnoreForPayment   bool
}

// Item represents an item
type Item struct {
	// Define fields as needed
}

// CheckIfLineOnHold checks if the line item is on hold
func (li *LineItem) CheckIfLineOnHold() bool {
	for _, hold := range li.HoldHistory {
		if hold.CheckIfHoldCreated() {
			return true
		}
	}
	return false
}

// CheckIfHoldCreated checks if hold is created in HoldHistory
func (hh *HoldHistory) CheckIfHoldCreated() bool {
	// Add logic to check if hold is created
	return false
}

// validate is a placeholder function for validation logic
func (li *LineItem) validate() error {
	// Add validation logic here
	return nil
}

func main() {
	// Example usage
	lineItem := LineItem{
		LineNumber: 1,
		// Initialize other fields as needed
	}

	err := lineItem.validate()
	if err != nil {
		// Handle validation error
	}

	onHold := lineItem.CheckIfLineOnHold()
	fmt.Println("Is LineItem on hold?", onHold)
}





package main

import "fmt"

// RewardsOrderStatus is an enum-like type
type RewardsOrderStatus string

// Note represents a note
type Note struct {
	// Define fields as needed
}

// HoldHistory represents hold history
type HoldHistory struct {
	// Define fields as needed
}

// LinePriceInfo represents line price information
type LinePriceInfo struct {
	// Define fields as needed
}

// OrderLineAdditionalInformation represents additional information for an order line
type OrderLineAdditionalInformation struct {
	// Define fields as needed
}

// LineItem represents a line item in an order
type LineItem struct {
	Item               Item
	OrderedQuantity    string
	Status             RewardsOrderStatus
	LockedForUpdate    *bool
	LineNumber         int `validate:"min=1,nonzero"`
	SubLineNumber      *int
	LinePriceInfo      LinePriceInfo
	Notes              []Note
	HoldHistory        []HoldHistory
	AdditionalInfo     OrderLineAdditionalInformation
	IgnoreForPayment   bool
}

// Item represents an item
type Item struct {
	// Define fields as needed
}

// CheckIfLineOnHold checks if the line item is on hold
func (li *LineItem) CheckIfLineOnHold() bool {
	return li.CheckIfHoldHistoryCreated()
}

// CheckIfHoldHistoryCreated checks if any hold history indicates a hold is created
func (li *LineItem) CheckIfHoldHistoryCreated() bool {
	for _, hold := range li.HoldHistory {
		if hold.CheckIfHoldCreated() {
			return true
		}
	}
	return false
}

// CheckIfHoldCreated checks if hold is created in HoldHistory
func (hh *HoldHistory) CheckIfHoldCreated() bool {
	// Add logic to check if hold is created
	return false
}

// validate is a placeholder function for validation logic
func (li *LineItem) validate() error {
	// Add validation logic here
	return nil
}

func main() {
	// Example usage
	lineItem := LineItem{
		LineNumber: 1,
		// Initialize other fields as needed
	}

	err := lineItem.validate()
	if err != nil {
		// Handle validation error
	}

	onHold := lineItem.CheckIfLineOnHold()
	fmt.Println("Is LineItem on hold?", onHold)
}
