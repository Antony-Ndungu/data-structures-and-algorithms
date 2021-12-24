const treeify = require("treeify");

class Node {
    constructor(value){
        this.value = value;
        this.left = null;
        this.right = null;
    }
}

class BST {
    constructor(){
        this.root = null;
    }

    insert(value){
        const newNode = new Node(value);

        if(this.root === null){
            this.root = newNode;
            return this;
        }

        let current = this.root;

        while(true){
            if(value < current.value){
                if(current.left === null){
                    current.left = newNode;
                    break;
                }
                current = current.left;
            } else if(value > current.value){
                if(current.right === null){
                    current.right = newNode;
                    break;
                }
                current = current.right;
            }else{
                break;
            }
        }

        return this;
    }

    contains(value){
        if(this.root === null) return false;

        let current = this.root;
        while(current !== null){
            if(value < current.value){
                current = current.left;
            } else if(value > current.value){
                current = current.right;
            } else{
                return true;
            }
        }

        return false;
    }

    remove(value){
        if(this.root === null) return;

        let current = this.root;
        let parent = null;
        while(current !== null){
            if(value < current.value){
                parent = current;
                current = current.left;
            } else if(value > current.value){
                parent = current;
                current = current.right;
            } else {
                if(current.left === null && current.right === null){
                    if(parent === null){
                        this.root = null;
                        return;
                    }

                    if(current === parent.left){
                        parent.left = null;
                        return;
                    }
                    parent.right = null;
                    return;
                } else if(current.left !== null && current.right === null){
                    if(parent === null){
                        this.root = current.left;
                        return;
                    }

                    if(current === parent.left){
                        parent.left = current.left;
                        return;
                    }

                    parent.right = current.left;
                    return;
                } else if(current.left === null && current.right !== null){
                    if(parent === null){
                        this.root = current.right;
                        return;
                    }

                    if(current === parent.left){
                        parent.left = current.right;
                        return;
                    }

                    parent.right = current.right;
                    return;
                } else {
                    let { predecessor, predecessorParent} = this.getPredecessor(current.left, current);
                    current.value = predecessor.value;
                    if(predecessorParent === null){
                        return;
                    }
                    if(predecessor === predecessorParent.left){
                        predecessorParent.left = null;
                        return;
                    }
                    predecessorParent.right = null;
                    return;
                }
            }
        }
    }

    getPredecessor(node, parent){
        let current = node;
        while(current.right !== null){
            parent = current;
            current = current.right;
        }
        return {predecessor: current, predecessorParent: parent};
    }
}

class AVLTree extends BST{
    constructor(){
        this.height = 0;
    }
}

const tree = new BST();
tree.insert(20).insert(25).insert(11).insert(30).insert(22).insert(21).insert(23);
console.log(treeify.asTree(tree, true));
tree.remove(22);
console.log(treeify.asTree(tree, true));