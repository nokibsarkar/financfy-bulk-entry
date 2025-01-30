"use client"
import { useCallback, useState } from "react";
const AddTransactionForm = ({type} : {type : string}) => {
    
    const [date, setDate] = useState('');
    const [voucherNo, setVoucherNo] = useState('');
    const [amount, setAmount] = useState('');
    const [contact, setContact] = useState('');
    const [reference, setReference] = useState('');
    const [remarks, setRemarks] = useState('');
    const [category, setCategory] = useState('');
    const [receiveMode, setReceiveMode] = useState('');
    const saveHandler = useCallback((e) => {
        const data = {
            date,
            voucherNo,
            amount,
            contact,
            reference,
            remarks,
            category,
            receiveMode
        };
        console.log(data);
    }, [])

    return (
      <form onSubmit={e => e.preventDefault()}>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-black">
              Date & time *
            </label>
            <input
              type="datetime-local"
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
              defaultValue="2"
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
              Receive mode *
            </label>
            <select className="w-full mt-1 p-2 border border-gray-300 rounded-md text-gray-700 text-xs"
                value={receiveMode}
                onChange={e => setReceiveMode(e.target.value)}
            >
              <option>Receive mode</option>
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
          >
            Save
          </button>
        </div>
      </form>
    )
  }
  export default AddTransactionForm;