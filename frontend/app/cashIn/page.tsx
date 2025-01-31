import ReturnButton from '@/components/ReturnButton';
import AddTransactionForm from '@/components/transaction/TransactionForm';
import { CashbookId } from '@/consts/server';
import React from 'react' 
export default function User() {
  return (      
         <main >
        
               <div className="bg-white p-6 flex justify-center text-black">
               
                  <div className="bg-white p-6 rounded-lg shadow-md w-full max-w-2xl">
                    <h2 className="text-lg text-primary mb-4">Add Cash in Entry</h2>
                    <AddTransactionForm type='cashIn' cashbookId={CashbookId} modeLabel='Receive mode' />
                  </div>


                </div>  
           
         </main>
        
  );
}
