import { useState } from "react"

const PatientForm = ({patients, setPatients, setSearchResults}) => {
    const [id, setId] = useState('0')
    const [ward, setWard] = useState('0')
    const [doctorId, setDoctorId] = useState('0')
    const [dischargeStatus, setDischargeStatus] = useState(false)

    const onAdd = (e) => {
        e.preventDefault()

        setPatients({
            ...patients,
            [id]: {id, ward, doctorId, dischargeStatus}
        })
    }

    const onDelete = (e) => {
       e.preventDefault()

       const ps = { ... patients }
       delete ps[id]
       setPatients(ps)
    }

    const onSearch = (e) => {
        e.preventDefault()

        setSearchResults({[id]: patients[id]})
    }

    return (
    <>
        <form style={{width: 500}} className="m-auto mt-5"> 
            <div className="mb-3">
                <label className="form-label">Patient Id</label>
                <input className="form-control" type="number" value={id} onChange={e => setId(e.target.value)}/>
            </div>

            <div className="mb-3">
                <label className="form-label">Ward Number</label>
                <input className="form-control" type="number" value={ward} onChange={e => setWard(e.target.value)}/>
            </div>

            <div className="mb-3">
                <label className="form-label">Doctor Id</label>
                <input className="form-control" type="number" value={doctorId} onChange={e => setDoctorId(e.target.value)}/>
            </div>

            <div className="mb-3 form-check">
                <input type="checkbox" className="form-check-input" checked={dischargeStatus} onChange={e => setDischargeStatus(e.target.checked)}/>
                <label className="form-check-label">Discharged</label>
            </div>

            <button className="btn btn-success ms-1" onClick={onAdd}>Add</button>
            <button className="btn btn-primary ms-1" onClick={onAdd}>Update</button>
            <button className="btn btn-danger ms-1" onClick={onDelete}>Delete</button>
            <button className="btn btn-info ms-1" onClick={onSearch}>Search</button>
        </form>


    </>)
}

export default PatientForm