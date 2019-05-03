import React, {Component} from 'react';
import './App.css';

import TempGraph from './components/Graph'
import SelectSensors from './components/SelectSensors'
import moment from 'moment';

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
            selectedSensor: null,
            tempData: null,
            refreshRate: 350000

        };
        this.setSelectedSensor = this.setSelectedSensor.bind(this)
        this.setRefreshRate = this.setRefreshRate.bind(this)
        this.resetInterval = this.resetInterval.bind(this)
    }

    getTimeRange() {
        return "/" + moment().subtract(5, 'minute').format('YYYY-MM-DDTHH:mm:ss') + "/" + moment().format('YYYY-MM-DDTHH:mm:ss')
    }

    componentDidMount() {
        this.resetInterval();
    }

    setSelectedSensor(sensor) {
        this.setState({
            selectedSensor: sensor,
            tempData: null,
        });
    }

    setRefreshRate(rate) {
        this.setState({
            refreshRate: rate,
        }, this.resetInterval);
    }

    resetInterval() {
        if (this.timer)
            clearInterval(this.timer);
        this.timer = setInterval(
            _ => {
                if (this.state.selectedSensor) {
                    fetch(temp_url + this.state.selectedSensor + this.getTimeRange())
                        .then(data => data.json())
                        .then(json => {
                            if (json)
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

                    fetch(mem_url + this.state.selectedSensor + this.getTimeRange())
                        .then(data => data.json())
                        .then(json => {
                            if (json)
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

                    fetch(cpu_url + this.state.selectedSensor + this.getTimeRange())
                        .then(data => data.json())
                        .then(json => {
                            if (json)
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
                }
            }, this.state.refreshRate)
        console.log(this.timer)
    }

    render() {
        return (
            <div className="App">
                <div style={styles.appBar}>
                    SpartanUp - Monitor
                </div>
                <SelectSensors
                    refreshRate={this.state.refreshRate}
                    setSelectedSensor={this.setSelectedSensor}
                    setRefreshRate={this.setRefreshRate}
                />
                <TempGraph data={this.state.tempData}/>
            </div>

        );
    }
}

export default App;
