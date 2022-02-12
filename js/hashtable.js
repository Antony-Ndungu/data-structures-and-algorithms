const hash = (key, size) => {
    let hashedKey = 37;
 
    for (let i = 0; i < Math.min(key.length, 100); i++) {
        hashedKey += key.charCodeAt(i);
    }
 
    return hashedKey % size;
 }

class HashTable {
    constructor (size=1023) {
        this.size = size;
        this.table = new Array(size);
    }

    set (key, value) {
        const index = hash(key, this.size);
        if (this.table[index] === undefined) {
            this.table[index] = [];
        }

        for (let pair of this.table[index]) {
            if (pair[0] === key) {
                pair[1] = value;
                return;
            }
        }
        this.table[index].push([key, value]);
    }

    get (key) {
        const index = hash(key, this.size);
        if (this.table[index] === undefined) return undefined;

        for (let pair of this.table[index]) {
            if (pair[0] === key) {
                return pair[1];
            }
        }

        return undefined;
    }

    keys () {
        let result = [];

        for (let bucket of this.table) {
            if (bucket === undefined) continue;
            for (let pair of bucket) {
                result.push(pair[0]);
            }
        }

        return result;
    }

    values () {
        let result = [];

        for (let bucket of this.table) {
            if (bucket === undefined) continue;
            for (let pair of bucket) {
                result.push(pair[1]);
            }
        }

        return result;
    }
}

const myHashTable = new HashTable();
myHashTable.set("firstname", "Antony");
myHashTable.set("lastname", "Maina");
myHashTable.set("lastname", "Ndungu");
console.log(myHashTable.get("firstname"));
console.log(myHashTable.get("lastname"));
console.log(myHashTable.get("middlename"));
console.log(myHashTable.keys());
console.log(myHashTable.values());



