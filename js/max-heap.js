class MaxHeap {
    constructor () {
        this.storage = [];
    }

    sort () {
        const result = [];

        while (this.storage.length > 0) {
            result.push(this.remove());
        }

        return result;
    }

    insert (data) {
        this.storage.push(data);
        this.heapifyUp();
    }

    heapifyUp () {
        let currentIndex = this.storage.length - 1;
        while (this.hasParent(currentIndex) === true) {
            const parentIndex = this.getParentIndex(currentIndex)
            if (this.getParent(currentIndex) < this.storage[currentIndex]) {
                this.swap(currentIndex, parentIndex);
            }
            currentIndex = parentIndex;
        }
    }

    remove () {
        if (this.storage.length === 0) return null;
        const maximum = this.storage[0];
        if (this.storage.length === 1) 
            this.storage.pop();
        else {
            this.storage[0] = this.storage.pop();
            this.heapifyDown();
        }
            
        return maximum;
    }

    heapifyDown() {
        let currentIndex = 0;
        while (this.hasLeftChild(currentIndex) === true) {
            let largestChildIndex = this.getLeftChildIndex(currentIndex);
            if (this.hasRightChild(currentIndex) && this.getRightChild(currentIndex) > this.storage[largestChildIndex])
                largestChildIndex = this.getRightChildIndex(currentIndex);
            if (this.storage[largestChildIndex] > this.storage[currentIndex]) {
                this.swap(largestChildIndex, currentIndex);
                currentIndex = largestChildIndex;
            } else {
                break;
            }
        }
    }

    swap (i, j) {
        const temp = this.storage[i];
        this.storage[i] = this.storage[j];
        this.storage[j] = temp;
    }

    getParent (childIndex) {
        return this.storage[this.getParentIndex(childIndex)];
    }

    getLeftChild (parentIndex) {
        return this.storage[this.getLeftChildIndex(parentIndex)];
    }

    getRightChild (parentIndex) {
        return this.storage[this.getRightChildIndex(parentIndex)];
    }

    hasParent (index) {
        return this.getParentIndex(index) >= 0;
    }

    hasLeftChild (index) {
        return this.getLeftChildIndex(index) < this.storage.length;
    }

    hasRightChild (index) {
        return this.getRightChildIndex(index) < this.storage.length;
    }
    
    getParentIndex (childIndex) {
        return Math.floor((childIndex - 1) / 2);
    }

    getLeftChildIndex (parentIndex) {
        return 2 * parentIndex + 1;
    }

    getRightChildIndex (parentIndex) {
        return 2 * parentIndex + 2;
    }
}
/*
                
*/

const myMaxHeap = new MaxHeap();
const arr = [10, 5, 1000, 50, 2, 8]
for (let element of arr) {
    myMaxHeap.insert(element);
}

console.log(myMaxHeap.sort());