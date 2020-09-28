package cli

import (
	"fmt"

	"github.com/via-justa/admiral/datastructs"
)

// CreateChildGroup accept parent and child groups to create new child-group relationship
func CreateChildGroup(parent datastructs.Group, child datastructs.Group) error {
	if child.ID == parent.ID {
		return fmt.Errorf("child and parent cannot be the same group")
	}

	b, err := isRelationshipLoop(child.ID, parent.ID)
	if err != nil {
		return err
	} else if b {
		return fmt.Errorf("relationship loop detected")
	}

	childGroup := datastructs.ChildGroup{
		Parent: parent.ID,
		Child:  child.ID,
	}

	// TODO: check parent is not child of child (relationship loop)
	i, err := db.insertChildGroup(childGroup)
	if err != nil {
		return err
	} else if i == 0 {
		return fmt.Errorf("no lines affected")
	}

	return nil
}

// ViewChildGroupsByParent accept parent group ID and return all child-group relationship for that group
func ViewChildGroupsByParent(parentID int) (childGroups []datastructs.ChildGroup, err error) {
	childGroups, err = db.selectChildGroup(0, parentID)
	if err != nil {
		return childGroups, err
	} else if childGroups == nil {
		return childGroups, fmt.Errorf("no record matched request")
	}

	return childGroups, nil
}

// ViewChildGroupsByChild accept child group ID and return all child-group relationship for that group
func ViewChildGroupsByChild(childID int) (childGroups []datastructs.ChildGroup, err error) {
	childGroups, err = db.selectChildGroup(childID, 0)
	if err != nil {
		return childGroups, err
	} else if childGroups == nil {
		return childGroups, fmt.Errorf("no record matched request")
	}

	return childGroups, nil
}

// ListChildGroups return all child-group relationships
func ListChildGroups() (childGroups []datastructs.ChildGroup, err error) {
	childGroups, err = db.getChildGroups()
	if err != nil {
		return childGroups, err
	} else if childGroups == nil {
		return childGroups, fmt.Errorf("no record matched request")
	}

	return childGroups, nil
}

// DeleteChildGroup accept child-group relationship to remove
func DeleteChildGroup(childGroup datastructs.ChildGroup) (affected int64, err error) {
	affected, err = db.deleteChildGroup(childGroup)
	if err != nil {
		return affected, err
	} else if affected == 0 {
		return affected, fmt.Errorf("no record matched")
	}

	return affected, nil
}

func getParents(child int, parents []int) ([]int, error) {
	childGroups, err := db.selectChildGroup(child, 0)
	if err != nil {
		return nil, err
	}

	for _, group := range childGroups {
		parents = append(parents, group.Parent)
		child = group.Parent

		parents, err = getParents(child, parents)
		if err != nil {
			return nil, err
		}
	}

	return parents, nil
}

func getChildren(parent int, children []int) ([]int, error) {
	parentGroups, err := db.selectChildGroup(0, parent)
	if err != nil {
		return nil, err
	}

	for _, group := range parentGroups {
		children = append(children, group.Child)
		parent = group.Child

		children, err = getChildren(parent, children)
		if err != nil {
			return nil, err
		}
	}

	return children, nil
}

func isRelationshipLoop(child, parent int) (bool, error) {
	children, err := getChildren(child, []int{})
	if err != nil {
		return false, err
	}

	for _, c := range children {
		if parent == c {
			return true, nil
		}
	}

	return false, nil
}