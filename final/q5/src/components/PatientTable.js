
const PatientTable = ({patients}) => {

    const mapPatients = () => Object.keys(patients).map(key => {
        const p = patients[key]
        return <li key={key}>Id: {p.id} - Ward Number: {p.ward} - Doctor Id: {p.doctorId} - Discharged: {p.dischargedStatus ? "true" : "false"}</li>
    });

    return (
        <ul>
            {mapPatients()}
        </ul>
    )
}

export default PatientTable