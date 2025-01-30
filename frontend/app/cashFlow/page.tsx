import { BaseURL } from "@/consts/server";
import { SingleCashFlowProperty } from "@/types/cashflow";
import { ResponseMultiple } from "@/types/response";
import React from "react";

const SingleCashFlowRow = (p : SingleCashFlowProperty ) => (
    <tr className='' >
        <td className=''>{p.date}</td>
        <td className=' '>{p.openingBalance}</td>
        <td className=' '>{p.totalIn}</td>
        <td className=''>{p.totalOut}</td>
        <td className=' '>{p.netBalance}</td>
        <td className=' '>{p.closingBalance}</td>
        <td>
            <button className="action-btn">...</button>
        </td>
    </tr>
)
export default async function CashFlow() {
    // const res = await fetch('url', {
    //     method : 'GET',

    // }).then(a => a.json());
    const response = await fetch(BaseURL + '/api/v1/cashbook', {
        method : 'GET',
         
    })
    
    const res : ResponseMultiple<SingleCashFlowProperty>  = await response.json();
    const cashFlow = res.data;

  return (
    <div className="p-6 h-screen bg-white ml-48">
      <h2 className="text-lg  text-primary mb-4">
        Mst Rukaiya Islam Tonni's cash book : Daily cash flow
      </h2>

      <div className="overflow-x-auto">
        <table className="w-full border-collapse border border-gray-300 text-xs ">
          <thead className="bg-gray-100 ">
            <tr className="text-left text-black">
              <th className="border border-gray-300 p-3">DATE</th>
              <th className="border border-gray-300 p-3 text-center">OPENING BALANCE</th>
              <th className="border border-gray-300 p-3 text-center">TOTAL IN</th>
              <th className="border border-gray-300 p-3 text-center">TOTAL OUT</th>
              <th className="border border-gray-300 p-3 text-center">NET BALANCE</th>
              <th className="border border-gray-300 p-3 text-center">CLOSING BALANCE</th>
              <th className="border border-gray-300 p-3 text-center">ACTION</th>
            </tr>
          </thead>
            <tbody className='border-collapse border'>
                        {cashFlow?.map((p) => <SingleCashFlowRow {...p}  key={p.date}/>)}
                        
          </tbody>
        </table>
      </div>
    </div>
  );
};
