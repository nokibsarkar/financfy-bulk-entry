import { BaseURL } from '@/consts/server';
import { SingleCashbookRowProperty } from '@/types/cashbook';
import { ResponseMultiple } from '@/types/response'
import React from 'react' 


const SingleCashbookRow = (p : SingleCashbookRowProperty ) => (
    <tr className='' id={p.id}>
        <td className='py-2'>{p.name}</td>
        <td className=' '>{p.openingBalance}</td>
        <td className=' '>{p.totalIn}</td>
        <td className=''>{p.totalOut}</td>
        <td className=' '>{p.netBalance}</td>
        <td>
            <button className="action-btn">...</button>
        </td>
    </tr>
)
export default async function User() {
    // const res = await fetch('url', {
    //     method : 'GET',

    // }).then(a => a.json());
    const res : ResponseMultiple<SingleCashbookRowProperty> = await fetch(BaseURL + '/api/v1/cashbook', {
        method : 'GET',
         
    }).then(a => a.json());


    const cashbooks = res.data;
  return (      
         <main >
             <div className="min-h-screen bg-gray-50 p-6 ml-55 pl-60">
               
                <div className="flex justify-between items-center mb-6 mt-10">
                      <div>
                         <h1 className="text-lg text-black ">Cash books overview</h1>
                         <p className="text-black text-xs">Lets see an overview of your cash books</p>
                      </div>
                      <div className="flex justify-between gap-4 text-sm">
                          <div className="dropdown">
                              <select id="timezone" className='text-black'>
                                  <option className='text-black'>Bangladesh Time</option>
                                  <option className='text-black'>Asia/Dhaka</option>
                              </select>
                          </div>

                          <div className="dropdown">
                              <select className='text-black' id="sort">
                                  <option className='text-black'>Ascending</option>
                                  <option className='text-black'>Descending</option>
                              </select>
                          </div>

                          <div className="dropdown">
                              <select className='text-black' id="name">
                                  <option className='text-black'>Name</option>
                                  <option className='text-black'>Net Balance</option>
                                  <option className='text-black'>Cash In</option>
                                  <option className='text-black'>Cash Out</option>
                                  <option className='text-black'>Created date</option>
                              </select>
                          </div>

                          <div className="dropdown">
                              <select className='text-black' id="time-range">
                                  <option className='text-black'>Last two months</option>
                                  <option className='text-black'>This month</option>
                                  <option className='text-black'>Last month</option>
                                  <option className='text-black'>Last three months</option>
                                  <option className='text-black'>All time</option>
                              </select>
                          </div>
            
                      </div>
                
                </div>
               
                <div className="overflow-x-auto">
        <table className="w-full border-collapse border border-gray-300 text-xs ">
          <thead className="bg-primary ">
            <tr className="text-left text-white ">
              <th className="border border-gray-300 p-3 text-center">CASH BOOK</th>
              <th className="border border-gray-300 p-3 text-center">OPENING BALANCE</th>
              <th className="border border-gray-300 p-3 text-center">TOTAL IN</th>
              <th className="border border-gray-300 p-3 text-center">TOTAL OUT</th>
              <th className="border border-gray-300 p-3 text-center">NET BALANCE</th>
              <th className="border border-gray-300 p-3 text-center">ACTION</th>
            </tr>
          </thead>
          <tbody className='border-collapse border text-black text-center '>
                        {cashbooks?.map((p) => <SingleCashbookRow {...p} key={p.id} />)}
                        </tbody>
        </table>
      </div>
            
                <div className="cashbook-switcher text-black mt-56">
                    <div className="cashbook-dropdown flex">
                        <select className='border-2 border-primary rounded-md p-2 text-xs'>
                            <option>Mst Rukaiya Islam Tonni's Cash Book</option>
                        </select>
                    
                    <button className="admin-badge px-4 bg-primary ml-4 text-white text-sm rounded-md">Admin</button>
                    </div>
                    <div className="info-message text-xs mt-2">
                        <span className="info-icon">ℹ️</span> 
                        You can switch your active cash book from here.
                    </div>
                  
                </div>


             </div>
      
         </main>
        
  );
}
