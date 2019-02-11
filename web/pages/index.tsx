import * as React from 'react'
import {Pane} from 'evergreen-ui'
import {TransactionList} from "../components/TransactionList";
import {MenuBar} from "../components/Sidebar";


export default () => {
    return (
        <Pane flexGrow={1} display={"flex"}>
            <MenuBar/>
            <Pane overflow="auto">
                <TransactionList/>
            </Pane>
        </Pane>
    )
}