const aesjs = require('aes-js');

const key = [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1];
const key128 = new Uint8Array(key);
const encrypted = 'U2FsdGVkX19jIyfUv0bkdP9Kcr1yP83XEyGrEZ0Pdn9aGKWYufmKcOozJtPnB7UyklqanvTIwMkm/V2+foQA3BGlbkAnEhZyEsbyzCW8RjuOZ9M/Ir1aOe+cO+gSIImXPjH8hE/qIlkXB0piOW53AJTO52BS4s7kMOat+RChiDDVlBgZL3ABzAHHb/SU+LC211AARjdooKxnIEA/n/B6PA==';
const encryptedBytes = aesjs.utils.hex.toBytes(encrypted);

var aesCtr = new aesjs.ModeOfOperation.ctr(key, new aesjs.Counter(5));
var decryptedBytes = aesCtr.decrypt(encryptedBytes);

// Convert our bytes back into text
console.log(aesjs.utils.utf8.fromBytes(decryptedBytes));
