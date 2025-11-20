
/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-11-20
 * @fileoverview: This file calculates cents 
 */

// INPUT
// asks for an input and converts the input into a proper value 
const input = prompt("Enter the amount of cents: ") || "0";
const totalCents = parseInt(input);

// PROCESS
let centsLeft = totalCents;

// calculates toonies 
var toonies = (centsLeft - (centsLeft % 200)) / 200;
centsLeft = centsLeft % 200;

// calculates loonies 
var loonies = (centsLeft - (centsLeft % 100)) / 100;
centsLeft = centsLeft % 100;

// calculates quarters 
var quarters = (centsLeft - (centsLeft % 25)) / 25;
centsLeft = centsLeft % 25;

// calculates dimes 
var dimes = (centsLeft - (centsLeft % 10)) / 10;
centsLeft = centsLeft % 10;

// calculates nickels 
var nickels = (centsLeft - (centsLeft % 5)) / 5;
centsLeft = centsLeft % 5;

// OUTPUT
console.log(`Your change is: ${toonies} toonies, ${loonies} dollar, ${quarters} quarters, ${dimes} dimes, ${nickels} nickels and ${centsLeft} cents`);
console.log("\nDone.")