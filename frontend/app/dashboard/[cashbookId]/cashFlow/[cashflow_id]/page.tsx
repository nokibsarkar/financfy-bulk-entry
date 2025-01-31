
import Server from "@/consts/server";
import { SingleTransactionProperty } from "@/types/transaction";
import React from "react";
const SingleTransactionRow = (p : SingleTransactionProperty) => (
    <tr className='text-center' id={p.date}>
        <td className='py-2'>{p.voucherNo}</td>
        <td className=''>{p.attachments}</td>
        <td className=' '>{p.category}</td>
        <td className=''>{p.contact}</td>
        <td className=' '>{p.amount}</td>
        <td className=' '>{p.verify}</td>
        <td className=' '>{p.remarks}</td>
        <td className=' '>{p.action}</td>
        <td>
            <button className="action-btn">...</button>
        </td>
    </tr>
)
export default async function Transaction({params} : {params : {cashflow_id : string}}){
    const {cashflow_id : id}= await params
    const res = await Server.getTransactions(id);
    if(res.error){
        return <div>{res.error}</div>
    }
    const transactions = res.data;
    return(
        <div className="p-6 h-screen bg-white text-xs ml-48">
           
        <div className="overflow-x-auto">
          <table className="w-full border-collapse border border-gray-300">
            <thead className="bg-primary text-white">
              <tr className="text-white">
                <th className="border border-gray-300 p-3 text-left">DATE </th>
                <th className="border border-gray-300 p-3 text-center">VOUCHER NO.</th>
                <th className="border border-gray-300 p-3 text-center">ATTACHMENTS</th>
                <th className="border border-gray-300 p-3 text-center">CATEGORY</th>
                <th className="border border-gray-300 p-3 text-center">CONTACT</th>
                <th className="border border-gray-300 p-3 text-center">AMOUNT</th>
                <th className="border border-gray-300 p-3 text-center">VERIFY</th>
                <th className="border border-gray-300 p-3 text-center">REMARKS</th>
                <th className="border border-gray-300 p-3 text-center">ACTION</th>
              </tr>
            </thead>
           
            <tbody className='border-collapse border text-black text-center '>
                {transactions?.map((p) => <SingleTransactionRow {...p} key={p.id} />)}
            </tbody>
            
          </table>
        </div>
      </div>
    )
}