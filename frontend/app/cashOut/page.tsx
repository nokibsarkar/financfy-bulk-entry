import React from 'react' 
export default function User() {
  return (      
    <main >
    <div className="bg-gray-100 p-6 flex justify-center text-black">
       <div className="bg-white p-6 rounded-lg shadow-md w-full max-w-2xl">
         <h2 className="text-lg text-primary mb-4">Add Cash out Entry</h2>

         <form>
           <div className="grid grid-cols-2 gap-4">
             <div>
               <label className="block text-sm font-medium text-black">
                 Date & time *
               </label>
               <input
                 type="datetime-local"
                 className="w-full mt-1 p-2 border border-gray-300 rounded-md text-xs text-gray-700"
               />
             </div>

             <div>
               <label className="block text-sm font-medium text-black">
                 Voucher number *
               </label>
               <input
                 type="text"
                 defaultValue="2"
                 className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
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
               />
             </div>

             <div>
               <label className="block text-sm font-medium text-black">
                 Contact
               </label>
               <select className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs">
                 <option>Contact</option>
                 <option>Mst. Rukaiya Islam Tonni</option>
               </select>
             </div>

             <div>
               <label className="block text-sm font-medium  text-black">
                 Payment mode *
               </label>
               <select className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs">
                 <option>Payment mode</option>
                 <option>Cash</option>
                 <option>Bank</option>
                 <option>Online method</option>
               </select>
             </div>

             <div>
               <label className="block text-sm font-medium  text-black">
                 Category *
               </label>
               <select className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs">
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
               className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
             />
           </div>

           <div className="mt-4">
             <label className="block text-sm font-medium  text-black">
               Remarks
             </label>
             <textarea
               className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
               rows={3}
             ></textarea>
           </div>

           
           <div className="mt-6">
             <button
               type="submit"
               className="w-full bg-primary text-white py-2 px-4 rounded-md "
             >
               Save
             </button>
           </div>
         </form>
       </div>
</div>  

</main>
        
  );
}
