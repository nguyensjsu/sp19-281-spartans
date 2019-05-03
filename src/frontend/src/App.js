import React, {Component} from 'react';
import './App.css';

import TempGraph from './components/Graph'
import SelectSensors from './components/SelectSensors'

const temp_url = "";
const mem_url = "";
const cpu_url = "";


const styles = {
    appBar: {
        padding: "2vh",
        background: "#5497F3",
        color: "white",
        fontSize: "2em",
        textAlign: "center",
        boxShadow: "0 2px 2px #aaa"
    }
};

class App extends Component {

    constructor(props) {
        super(props);
        this.state = {
            selectedSensors: [],
            tempData: null,
            cpuData: null,
            memData: null,
        };
        this.setSelectedSensors = this.setSelectedSensors.bind(this)
        this.timer;
    }

    componentDidMount() {
        this.timer = setInterval(
            _ => {
                fetch(temp_url + "12345/2019-05-02T20:10:40/2019-05-02T22:50:40")
                    .then(data => data.json())
                    .then(json => {
                        this.setState({
                            tempData: {
                                x: 'x',
                                xFormat: '%Y',
                                columns: [
                                    ['x', ...json.map(item => item.Time)],
                                    ['Temp', ...json.map(item => item.Temperature)]
                                ],
                            }
                        })
                    });

                fetch(mem_url + "12345/2019-05-02T20:10:40/2019-05-02T22:50:40")
                    .then(data => data.json())
                    .then(json => {
                        this.setState({
                            tempData: {
                                x: 'x',
                                xFormat: '%Y',
                                columns: [
                                    ['x', ...json.map(item => item.Time)],
                                    ['Memory usage', ...json.map(item => item.Temperature)]
                                ],
                            }
                        })
                    });

                fetch(cpu_url + "12345/2019-05-02T20:10:40/2019-05-02T22:50:40")
                    .then(data => data.json())
                    .then(json => {
                        this.setState({
                            tempData: {
                                x: 'x',
                                xFormat: '%Y',
                                columns: [
                                    ['x', ...json.map(item => item.Time)],
                                    ['CPU Usage', ...json.map(item => item.Temperature)]
                                ],
                            }
                        })
                    });
            }, 1000)
    }

    setSelectedSensors(sensors) {
        this.setState({selectedSensors: sensors});
        console.log(sensors)
    }

    updateGraph(data) {
        console.log(data)
    }

    render() {
        return (
            <div className="App">
                <div style={styles.appBar}>
                    SpartanUp - Monitor
                </div>
                <SelectSensors setSelectedSensors={this.setSelectedSensors}/>
                <TempGraph data={this.state.tempData}/>
            </div>

        );
    }
}

export default App;
