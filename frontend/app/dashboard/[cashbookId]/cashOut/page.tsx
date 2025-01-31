import ReturnButton from '@/components/ReturnButton';
import AddTransactionForm from '@/components/transaction/TransactionForm';
import React from 'react'
export default async function CashOut({ params }: { params: { cashbookId: string } }) {
	const { cashbookId: CashbookId } = await params
	return (
		<main>
			<div className="bg-white flex p-6 justify-center text-black">
				<div className="bg-white p-6 rounded-lg shadow-md w-full max-w-2xl">
					<h2 className="text-lg text-primary mb-4">Add Cash out Entry</h2>
					<AddTransactionForm type='cashOut' cashbookId={CashbookId} modeLabel='Payment mode' />
				</div>
			</div>
		</main>
	);
}
