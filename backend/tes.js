const res = await fetch("https://f78f-103-230-40-62.ngrok-free.app/api/v1/cashbook", {
    "headers": {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:134.0) Gecko/20100101 Firefox/134.0",
        "Accept": "*/*",
        "Accept-Language": "en-US,en;q=0.5",
        "Content-Type": "text/plain;charset=UTF-8"
    },
    "method": "POST",
    "body": JSON.stringify({
        "date": "2022-06-30",
        "amount": 1000,
        "description": "Test"
    })
});
const data = await res.text();
console.log(data);