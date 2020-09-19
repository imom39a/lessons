package main

import "fmt"

func main(){

/*
        100
    90        110  
80        101     102
*/  
  tree := BinaryTree{
    rootNode: &BinaryTreeNode{
      Data: 100,
      left: &BinaryTreeNode{
        Data: 90, 
        left: &BinaryTreeNode{
          Data: 80,          
          left: &BinaryTreeNode{
            Data: 70,
            left: &BinaryTreeNode{
              Data: 60,          
              left: &BinaryTreeNode{
                Data: 50,          
              },
            },
          },
        },
      },
      right: &BinaryTreeNode{
        Data: 110,
        left: &BinaryTreeNode{
          Data: 101,
          },
        right: &BinaryTreeNode{
          Data: 112,
          },
        },
    },
  }

  fmt.Println(tree.IsSuperBalanced())  
}
