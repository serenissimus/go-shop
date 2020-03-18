import React from 'react';
import {Fragment, useEffect, useState} from 'react';
import axios from 'axios';
import { Technology } from './Technology';

export const App = () => {
    const [data, setData] = useState([]);
    const fetchData = async () => {
        const result = await axios(
          `${process.env.API_URL}/technologies`,
        );
        setData(result.data);
    };
    useEffect(() => { fetchData() }, []);

    return (<Fragment>
        <h1>Go Shop</h1>
        {data.map(it => (
            <Technology key={it.id} name={it.name} details={it.details}></Technology>
        ))}
    </Fragment>);
}
