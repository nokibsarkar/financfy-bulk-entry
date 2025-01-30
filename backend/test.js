const baseURL = "https://f78f-103-230-40-62.ngrok-free.app/api/v1";
async function NewCashbook() {
    const res = await fetch(baseURL + "/cashbook", {
        "method": "POST",
        "body": JSON.stringify({
            "name": "Cashbook",
            "userId": 67,
        })
    });
    const data = await res.json();
    console.log(data);
}
async function newTransaction() {
    const data = {
        "voucherNo": "123",
        "date": "2021-10-10",
        "category": "Income",
        "type" : "Cashout",
        "amount": 1000,
        "description": "Salary",
        "cashbookId": 1
    }
    const res = await fetch(baseURL + "/transaction", {
        "method": "POST",
        "body": JSON.stringify(data)
    });
    const resturnedData = await res.json();
    console.log(resturnedData);
}
async function newTransactions() {
    const data = [{
        "voucherNo": "123",
        "date": "2021-10-10",
        "category": "Income",
        "type" : "Cashout",
        "amount": 1000,
        "description": "Salary",
        "cashbookId": 1
    }]
    const res = await fetch(baseURL + "/transaction/bulk", {
        "method": "POST",
        "body": JSON.stringify(data)
    });
    const resturnedData = await res.json();
    console.log(resturnedData);
}
await NewCashbook();
await newTransaction();
await newTransactions();
