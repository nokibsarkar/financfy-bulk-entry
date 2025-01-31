import { SingleCashbookRowProperty } from "@/types/cashbook";
import { SingleCashFlowProperty } from "@/types/cashflow";
import { ResponseMultiple } from "@/types/response";
import { SingleTransactionProperty } from "@/types/transaction";

export const BaseURL = 'https://6fdd-118-179-192-216.ngrok-free.app';


class Server {
    static BaseURL = 'https://6fdd-118-179-192-216.ngrok-free.app';
    static async getTransactions(cashflow_id: string) : Promise<ResponseMultiple<SingleTransactionProperty>> {
        const qs = new URLSearchParams({
            cashflow_id
        }).toString()
        const res = await fetch(Server.BaseURL + '/api/v1/transaction/?' + qs, {
            method: 'GET'
        })
        return res.json();
    }
    static async getCashflows() : Promise<ResponseMultiple<SingleCashFlowProperty>> {
        const res = await fetch(Server.BaseURL + '/api/v1/cashflow/', {
            method: 'GET',
        })
        return res.json();
    }
    static async getSingleCashflow(id: string) : Promise<ResponseMultiple<SingleCashFlowProperty>> {
        const res = await fetch(Server.BaseURL + '/api/v1/cashflow/' + id, {
            method: 'GET',
        })
        return res.json();
    }
    static async getCashbooks() : Promise<ResponseMultiple<SingleCashbookRowProperty>> {
        const res = await fetch(Server.BaseURL + '/api/v1/cashbook/', {
            method: 'GET',
        })
        return res.json();
    }
    static async addBulkTransactions(data: SingleTransactionProperty[]) {
        const res = await fetch(BaseURL + "/api/v1/transaction/bulk", {
            "method": "POST",
            "body": JSON.stringify(data)
        });
        return res.json();
    }
}
export default Server;