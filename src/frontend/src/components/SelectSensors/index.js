import React, {Component} from 'react';
import Card from '@material-ui/core/Card';
import RadioGroup from '@material-ui/core/RadioGroup';
import Radio from '@material-ui/core/Radio';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';


const styles = {
    card: {
        padding: "2vh",
        width: "70vw",
        margin: '100px auto '
    }
};

class SelectSensors extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: []
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleRateChange = this.handleRateChange.bind(this);
    }

    componentDidMount() {
        fetch("")
            .then(data => data.json())
            .then(json => this.setState({data: json.map(({producer}) => ({id: producer, checked: false}))}))
    }

    handleChange(e) {
        if (this.props.setSelectedSensor)
            this.props.setSelectedSensor(e.target.value)

    }

    handleRateChange(e) {
        if (this.props.setRefreshRate)
            this.props.setRefreshRate(e.target.value)

    }

    render() {
        let {data} = this.state;
        return (
            <div>
                <Card style={styles.card}>
                    Device Selection:
                    <form>
                        <RadioGroup
                            name="sensor"
                            onChange={this.handleChange}>
                            {
                                data.map((id, index) =>
                                    <FormControlLabel
                                        key={index}
                                        control={<Radio/>}
                                        label={id.id + ""}
                                        value={id.id + ""}
                                    />
                                )
                            }
                        </RadioGroup>
                    </form>
                    Refresh Rate:

                    <FormControl>
                        <Select
                            value={this.props.refreshRate}
                            onChange={this.handleRateChange}>
                            <MenuItem value={1000}>Every Second</MenuItem>
                            <MenuItem value={10000}>Every 10 Second</MenuItem>
                            <MenuItem value={350000}>5 Min</MenuItem>
                        </Select>
                    </FormControl>


                </Card>
            </div>
        );
    }
}

export default SelectSensors;
