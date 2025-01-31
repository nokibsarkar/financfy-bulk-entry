'use client'
import Server from '@/consts/server';
import { SingleTransactionProperty } from '@/types/transaction';
import React, { useEffect } from 'react'

const BulkEntryRow = ({ removeRowByID, id, dispatchChange, transaction }: any) => {
    
    return (
        <tr className="border-b text-xs" id={id}>
            <td className="p-2">1</td>
            <td className="p-2">
                <input type="number"
                    className="w-full p-1 border rounded"
                    value={transaction.amount}
                    onChange={e => dispatchChange('amount', e.target.value)}
                />
            </td>
            <td className="p-2"><select className="w-full p-1 border rounded" value={transaction.type} onChange={e => dispatchChange('type', e.target.value)}>
                <option>Select...</option>
                <option>Cash In</option>
                <option>Cash Out</option>
            </select></td>
            <td className="p-2"><select className="w-full p-1 border rounded" value={transaction.contact} onChange={e => dispatchChange('contact', e.target.value)}>
                <option>Select...</option>
                <option>Mst Rukaiya Islam Tonni</option>
            </select></td>
            <td className="p-2"><select className="w-full p-1 border rounded" value={transaction.category} onChange={e => dispatchChange('category', e.target.value)}>
                <option>Select...</option>
                <option>Salary</option>
            </select></td>
            <td className="p-2"><select className="w-full p-1 border rounded" value={transaction.mode} onChange={e => dispatchChange('mode', e.target.value)}>
                <option disabled>Select...</option>
                <option>Cash</option>
                <option>Online method</option>
                <option>Bank</option>
            </select></td>
            <td className="p-2"><input type="text" className="w-full p-1 border rounded" placeholder="Enter remarks" value={transaction.remarks} onChange={e => dispatchChange('remarks', e.target.value)} /></td>
            <td className="p-2 text-center">
                <button className="text-red-500" onClick={removeRowByID}>Remove</button>
            </td>
        </tr>
    )
}
export default function User() {
    const [rows, setRows] = React.useState<React.ReactNode[]>([]);
    const [refs, setRefs] = React.useState<{ [k: string]: { current: any } }>({});
    const cashbookID = '1rj4m1je40000';
    const [tableRowMap, setTableRowMap] = React.useState<{ [k: string]: SingleTransactionProperty }>({});
    const [date, setDate] = React.useState('');
    const [voucherNo, setVoucherNo] = React.useState('');
    const removeRowByID = (id: string) => () => {
        const ref = refs[id];
        if (ref && ref.current)
            delete refs[id];
        setRows(rows.filter((r: any) => r.id !== id));
    }
    const dispatchChange = (id: string) => (key: string, value: string) => {
        setTableRowMap({ ...tableRowMap, [id]: { ...tableRowMap[id], [key]: value } });
    }
    const addRow = () => {
        const uniqueID = Math.random().toString(36).substring(7);
        const ref = React.createRef();
        setRefs({ ...refs, [uniqueID]: ref });
        setTableRowMap({
            ...tableRowMap, [uniqueID]: {
                id: uniqueID,
                amount: '',
                contact: '',
                remarks: '',
                category: '',
                type: '',
                mode: '',
            }
        });
        setRows([...rows, <BulkEntryRow ref={ref} removeRowByID={removeRowByID(uniqueID)} id={uniqueID} key={uniqueID} />]);
    }
    const saveBulkTransaction = () => {
        const data : SingleTransactionProperty[] = Object.values(tableRowMap).length ? Object.values(tableRowMap).map(r => {
            const { amount, contact, remarks, category, type, mode } = r;
            return {
                id : Math.random().toString(36).substring(7),
                amount : parseFloat(amount),
                contact,
                remarks,
                category,
                type,
                mode,
                date : new Date(date || '').toISOString(),
                voucherNo : voucherNo,
                cashbookId : cashbookID
            }
        }) : [];
        Server.addBulkTransactions(data).then(res => {
            console.log(res);
        }).catch(err => {
            console.log(err);
        }).finally(() => {
            console.log('finally');
        })
    }
    useEffect(() => {
        addRow();
    }, []);
    return (
        <main className='text-black bg-white ml-48' >
            <div className=" w-full h-screen bg-white p-6 ">
                <h2 className="text-lg mb-4">Add Bulk Transaction Entries</h2>

                <div className="grid grid-cols-2 gap-4 mb-4 text-xs">
                    <div>
                        <label className="block text-sm font-medium">Date & Time *</label>
                        <input type="date" className="mt-1 w-full p-2 border rounded" value={date} onChange={e => setDate(e.target.value)} />
                    </div>
                    <div>
                        <label className="block text-sm font-medium">Voucher Number *</label>
                        <input type="number" className="mt-1 w-full p-2 border rounded" value={voucherNo} onChange={e => setVoucherNo(e.target.value)} />
                    </div>
                </div>

                <table className="w-full border-collapse border rounded-lg overflow-hidden">
                    <thead>
                        <tr className="bg-gray-200 text-xs">
                            <th className="p-2">#</th>
                            <th className="p-2">Amount *</th>
                            <th className="p-2">Type *</th>
                            <th className="p-2">Contact</th>
                            <th className="p-2">Category *</th>
                            <th className="p-2">Payment Mode *</th>
                            <th className="p-2">Remarks</th>
                            <th className="p-2">Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {/* {rows} */}
                        {Object.values(tableRowMap)
                            .map((r, i) => <BulkEntryRow
                                ref={refs[r.id]}
                                removeRowByID={removeRowByID(r.id)}
                                id={r.id} key={r.id}
                                dispatchChange={dispatchChange(r.id)}
                                transaction={r}
                            />)}
                    </tbody>
                </table>

                <div className="mt-4 flex justify-end gap-4 text-xs">
                    <button className="px-4 py-2 bg-red-500 text-white rounded">Reset</button>
                    <button className="px-4 py-2 bg-blue-500 text-white rounded" onClick={saveBulkTransaction}>Save</button>
                    <button className="px-4 py-2 bg-green-500 text-white rounded" onClick={addRow}>Add Row</button>

                </div>
            </div>

        </main>

    );
}
