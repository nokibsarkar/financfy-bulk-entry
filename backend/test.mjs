const baseURL = "http://localhost:8080/api/v1";
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
        "cashbookId": 2147010522703200300
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
        "type" : "cashin",
        "amount": 1000,
        "description": "Salary",
        "cashbookId": 2147010522703200300
    }, {
        "voucherNo": "123",
        "date": "2021-10-10",
        "category": "Income",
        "type" : "cashin",
        "amount": 1000,
        "description": "Salary",
        "cashbookId": 2147010522703200300
    }, {
        "voucherNo": "123",
        "date": "2021-10-10",
        "category": "Income",
        "type" : "cashin",
        "amount": 1000,
        "description": "Salary",
        "cashbookId": 2147010522703200300
    }
        ];
    const res = await fetch(baseURL + "/transaction/bulk", {
        "method": "POST",
        "body": JSON.stringify(data)
    });
    const resturnedData = await res.json();
    console.log(resturnedData);
}

var k = true;
const millis = 10000;
setTimeout(() => {
    k = false;
}, millis);
for (var i = 0; i < 1000000; i++) {
    if (!k) {
        break;
    }
    await newTransactions();
    await NewCashbook();
    await newTransaction();
}
console.log(`Inserted ${i * 1} transactions in ${millis} milliseconds`);
// 60 transactions per second
// 30 cashbooks per second
