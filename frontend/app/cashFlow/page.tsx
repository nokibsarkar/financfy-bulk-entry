import Server, { BaseURL, CashbookId } from "@/consts/server";
import { SingleCashFlowProperty } from "@/types/cashflow";
import { ResponseMultiple } from "@/types/response";
import Link from "next/link";
import React from "react";

const SingleCashFlowRow = (p: SingleCashFlowProperty) => (
	<tr className='text-black text-center' >
		<td className='text-black p-3'>
			<Link href={`cashFlow/${p.id}`} className="text-primary">{p.date.split('T')[0]}</Link>
		</td>
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
	const res: ResponseMultiple<SingleCashFlowProperty> = await Server.getCashflows(CashbookId);
	const cashFlow = res.data;
	return (
		<div className="p-6 h-screen bg-white ml-48">
			<h2 className="text-lg  text-primary mb-4">
				Mst Rukaiya Islam Tonni's cash book : Daily cash flow
			</h2>
			<div className="overflow-x-auto">
				<table className="w-full border-collapse border border-gray-300 text-xs ">
					<thead className="bg-primary  ">
						<tr className="text-left text-white">
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
						{cashFlow?.map((p) => <SingleCashFlowRow {...p} key={p.id} />)}
					</tbody>
				</table>
			</div>
		</div>
	);
};
