const url = 'https://httpbin.org/get';
const method = 'GET';
const headers = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
}
const options = {};
const response = await fetch(url, options);
console.log(await response.json())