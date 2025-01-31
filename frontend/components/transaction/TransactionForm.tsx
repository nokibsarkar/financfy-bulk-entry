"use client"
import { BaseURL } from "@/consts/server";
import { ResponseMultiple } from "@/types/response";
import { useState } from "react";
type TransactionType = 'cashIn' | 'cashOut';
type AddTransactionFormProps = {
	type: string;
	date: string;
	voucherNo: string;
	amount: number;
	contact: string;
	reference: string;
	remarks: string;
	category: string;
	receiveMode: string;
	cashbookId: string;
};
const AddTransactionForm = ({ type, cashbookId = '' , modeLabel = 'Mode'}: { type: TransactionType, cashbookId: string, modeLabel : string }) => {
	const [date, setDate] = useState(new Date().toISOString().split('T')[0]);
	const [voucherNo, setVoucherNo] = useState('');
	const [amount, setAmount] = useState('');
	const [contact, setContact] = useState('');
	const [reference, setReference] = useState('');
	const [remarks, setRemarks] = useState<string>('');
	const [category, setCategory] = useState<string>('');
	const [receiveMode, setReceiveMode] = useState('');
	const [loading, setLoading] = useState<boolean>(false);
	const saveHandler = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
		const data: AddTransactionFormProps = {
			"voucherNo": voucherNo,
			"date": new Date(date).toISOString(),
			"category": category,
			"type": type,
			"amount": Number(amount),
			"cashbookId": cashbookId,
			"receiveMode": receiveMode,
			"contact": contact,
			"reference": reference,
			"remarks": remarks
		}
		const url = BaseURL + '/api/v1/transaction/';
		setLoading(true);
		fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		}).then(a => a.json()).then((a: ResponseMultiple<{ [k: string]: any }>) => {
			if (!a.error) {
				alert('Transaction saved successfully');
			} else {
				throw new Error(a.error);
			}
		}).catch(a => {
			console.log(a);
		}).finally(() => {
			setLoading(false);
		});
	};

	return (
		<form onSubmit={e => e.preventDefault()}>
			<div className="grid grid-cols-2 gap-4">
				<div>
					<label className="block text-sm font-medium text-black">
						Date *
					</label>
					<input
						type="date"
						className="w-full mt-1 p-2 border border-gray-300 rounded-md text-xs text-gray-700"
						value={date}
						onChange={e => setDate(e.target.value)}
					/>
				</div>

				<div>
					<label className="block text-sm font-medium text-black">
						Voucher number *
					</label>
					<input
						type="text"
						className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
						value={voucherNo}
						onChange={e => setVoucherNo(e.target.value)}
					/>
				</div>

				<div>
					<label className="block text-sm font-medium text-black">
						Amount *
					</label>
					<input
						type="text"
						placeholder="৳ 0.00"
						className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
						value={amount}
						onChange={e => setAmount(e.target.value)}
					/>
				</div>

				<div>
					<label className="block text-sm font-medium text-black">
						Contact
					</label>
					<select className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
						value={contact}
						onChange={e => setContact(e.target.value)}
					>
						<option>Contact</option>
						<option>Mst. Rukaiya Islam Tonni</option>
					</select>
				</div>

				<div>
					<label className="block text-sm font-medium  text-black">
						{modeLabel} *
					</label>
					<select className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
						value={receiveMode}
						onChange={e => setReceiveMode(e.target.value)}
					>
						<option disabled>Receive mode</option>
						<option>Cash</option>
						<option>Bank</option>
						<option>Online method</option>
					</select>
				</div>

				<div>
					<label className="block text-sm font-medium  text-black">
						Category *
					</label>
					<select
						value={category}
						onChange={e => setCategory(e.target.value)}
						className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs">
						<option>Category</option>
						<option>Salary</option>
					</select>
				</div>
			</div>

			<div className="mt-4">
				<label className="block text-sm font-medium  text-black">
					Reference
				</label>
				<input
					type="text"
					value={reference}
					onChange={e => setReference(e.target.value)}
					className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
				/>
			</div>

			<div className="mt-4">
				<label className="block text-sm font-medium  text-black">
					Remarks
				</label>
				<textarea
					value={remarks}
					onChange={e => setRemarks(e.target.value)}
					className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
					rows={3}
				></textarea>
			</div>


			<div className="mt-6">
				<button
					type="submit"
					className="w-full bg-primary text-white py-2 px-4 rounded-md "
					onClick={saveHandler}
					disabled={loading}
				>
					Save
				</button>
			</div>
		</form>
	)
}
export default AddTransactionForm;