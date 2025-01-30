import React from 'react' 
export default function User() {
  return (      
    <div className="p-6 h-screen bg-white text-xs ml-48">
    <div className="overflow-x-auto">
      <table className="w-full border-collapse border border-gray-300">
        <thead className="bg-gray-100">
          <tr className="text-gray-700">
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
        <tbody>
          
        </tbody>
      </table>
    </div>
  </div>
  );
}