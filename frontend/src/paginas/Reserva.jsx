import React, {useState} from "react";
import DatePicker from "react-date-picker";

const Reserva = () => {
    const DatePickerExample = () => {
        const [startDate, setStartDate] = useState(new Date());
        const [endDate, setEndDate] = useState(new Date());

        return (
            <div>
                <DatePicker
                    value={startDate}
                    onChange={date => setStartDate(date)}
                />
                <DatePicker
                    value={endDate}
                    onChange={date => setEndDate(date)}
                />
            </div>
        );
    };
}
export default Reserva;