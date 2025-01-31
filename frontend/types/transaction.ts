export type SingleTransactionProperty = {
    id : string
    date : string
    voucherNo : number
    attachments : number
    category : string
    contact : number
    amount: number
    verify : number
    remarks: string
    action : string
}
export type SingleTransactionInput = {
    id: string;
    amount: number | string;
    contact: string;
    remarks: string;
    category: string;
    type: string;
    mode: string;
}
