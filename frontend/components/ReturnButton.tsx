"use client"
const ReturnButton = ({to} : {to?: string}) => {
    console.log(to)
    return (
        <button onClick={() => window.history.back()} className="text-white text-xs bg-primary p-2  rounded-md">Go back</button>
    )
}
export default ReturnButton;