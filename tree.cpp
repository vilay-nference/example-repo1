#include<iostream>

class TreeNode {
    private:
    int data;
    TreeNode *left;
    TreeNode *right;
    
    public:
    TreeNode(int data);
    TreeNode* getLeft();
    TreeNode* getRight();
    void setLeft(TreeNode * left);
    void setRight(TreeNode * right);
    void printPreorder();
};

TreeNode::TreeNode(int data){
    this->data = data;
    this->left = NULL;
    this->right = NULL;
}

TreeNode* TreeNode::getLeft() {
    return this->left;
}

void TreeNode::setLeft(TreeNode* left) {
    this->left = left;
}

TreeNode* TreeNode::getRight() {
    return this->right;
}

void TreeNode::setRight(TreeNode* right) {
    this->right = right;
}

void TreeNode::printPreorder() {
    std::cout<<this->data<< "\t";
    if(this->left != NULL)
        this->left->printPreorder();
    if(this->right != NULL)
        this->right->printPreorder();
}

class Tree {
    private:
    TreeNode *root;
    public:
    Tree(TreeNode* root) {
        this->root = root;
    }
    void printTree();
    void addTree();
};

void Tree::printTree() {
    std::cout<< "Started printing tree" << std::endl;
    this->root->printPreorder();
}

void Tree::addTree() {

}

int main() {
    TreeNode* root = new TreeNode(1);
    TreeNode* l = new TreeNode(2);
    TreeNode* r = new TreeNode(3);
    TreeNode* lr = new TreeNode(4);

    root->setLeft(l);
    root->setRight(r);
    l->setRight(lr);

    Tree tree(root);
    tree.printTree();

    return 0;
}
