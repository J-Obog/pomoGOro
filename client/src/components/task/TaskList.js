import { React, useState, useContext, useEffect } from 'react'; 
import { AuthContext } from '../../context/AuthContext';
import axios from 'axios'; 
import Task from './Task'; 

const TaskList = () => {
    const { access } = useContext(AuthContext); 
    const [tasks, setTasks] = useState([]); 

    useEffect(() => {
        axios.get(`${process.env.REACT_APP_API_URL}/task/`, { headers: { 'Authorization': access }})
        .then(({data}) => {
            setTasks(data); 
        })
        .catch(err => {})
    }, [access])

    const addTask = async (e) => {
        const title = e.target.value;

        if(e.key === 'Enter') {
            try {
                const { data } = await axios.post(`${process.env.REACT_APP_API_URL}/task/`, {title: title}, { headers: { 'Authorization': access }});
                const newList = tasks.concat(data);
                setTasks(newList);
                e.target.value = ''; 
            } catch(err) {

            }
        }
    }

    return (
        <div className="bg-transparent flex flex-col justify-around">
            <div className="bg-black text-white text-lg rounded-lg p-2 mb-8">
                <h1>Tasks</h1>
            </div>
            <div className="mb-10">
                <input onKeyPress={addTask} 
                    placeholder="New Task" className="w-full outline-none border-b-2"/>
            </div>
            <div>
                { tasks.map(task => ( <Task key={task.id} {...task}/> )) }
            </div>
        </div>
    )
    
}

export default TaskList; 