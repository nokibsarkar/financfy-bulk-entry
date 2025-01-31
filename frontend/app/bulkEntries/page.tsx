'use client'
import Server, { CashbookId } from '@/consts/server';
import { SingleTransactionInput } from '@/types/transaction';
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
                <option disabled>Select...</option>
                <option value="cashIn">Cash In</option>
                <option value="cashOut">Cash Out</option>
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
                <option value="cash">Cash</option>
                <option value="online">Online method</option>
                <option value="bank">Bank</option>
            </select></td>
            <td className="p-2"><input type="text" className="w-full p-1 border rounded" placeholder="Enter remarks" value={transaction.remarks} onChange={e => dispatchChange('remarks', e.target.value)} /></td>
            <td className="p-2 text-center">
                <button className="text-red-500" onClick={removeRowByID}>Remove</button>
            </td>
        </tr>
    )
}
export default function BulkEntry() {
    const [refs, setRefs] = React.useState<{ [k: string]: { current: any } }>({});
    const [tableRowMap, setTableRowMap] = React.useState<{ [k: string]: SingleTransactionInput }>({});
    const [date, setDate] = React.useState('');
    const [voucherNo, setVoucherNo] = React.useState('');
    const removeRowByID = (id: string) => () => {
        setTableRowMap(Object.keys(tableRowMap).reduce((acc, k) => {
            if (k !== id) {
                acc[k] = tableRowMap[k];
            }
            return acc;
        }, {} as { [k: string]: SingleTransactionInput }));
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
                type: 'cashIn',
                mode: 'cash',
            }
        });
        
    }
    const saveBulkTransaction = () => {
        try {
            const dateParsed = new Date(date || '').toISOString();
            const data: SingleTransactionInput[] = Object.values(tableRowMap).length ? Object.values(tableRowMap).map(r => {
                const { amount, contact, remarks, category, type, mode } = r;
                return {
                    id: Math.random().toString(36).substring(7),
                    amount: parseFloat(amount as string),
                    contact,
                    remarks,
                    category,
                    type,
                    mode,
                    date: dateParsed,
                    voucherNo: voucherNo,
                    cashbookId: CashbookId
                }
            }) : [];
            Server.addBulkTransactions(data).then(res => {
                console.log(res);
                alert('Transaction saved successfully');
            }).catch(err => {
                console.log(err);
            }).finally(() => {
                console.log('finally');
            })
        } catch (e) {
            console.log('error parsing date', e);
        }
    }
    useEffect(() => {
        reset();
    }, []);
    const reset = () => {
        setDate(new Date().toISOString().split('T')[0]);
        setVoucherNo('');
        setTableRowMap({});
        setRefs({});
        addRow();
    }
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

                <table className="w-full border-collapse border  overflow-hidden">
                    <thead>
                        <tr className="bg-primary text-white border-collapse text-xs">
                            <th className="p-2 border">#</th>
                            <th className="p-2 border">Amount *</th>
                            <th className="p-2 border">Type *</th>
                            <th className="p-2 border">Contact</th>
                            <th className="p-2 border">Category *</th>
                            <th className="p-2 border">Payment Mode *</th>
                            <th className="p-2 border">Remarks</th>
                            <th className="p-2 border">Action</th>
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
                    <button className="px-4 py-2 bg-red-500 text-white rounded" onClick={reset}>Reset</button>
                    <button className="px-4 py-2 bg-blue-500 text-white rounded" onClick={saveBulkTransaction}>Save</button>
                    <button className="px-4 py-2 bg-green-500 text-white rounded" onClick={addRow}>Add Row</button>
                </div>
            </div>

        </main>

    );
}
