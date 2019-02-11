import * as React from 'react'
import {Pane, Heading, majorScale} from 'evergreen-ui'
import {MenuBar} from "../components/Sidebar";


export default () => {
    return (
        <Pane flexGrow={1} display={"flex"} >
            <MenuBar/>
            <Pane overflow="auto" paddingX={majorScale(5)} border margin={majorScale(3)} backgroundColor="var(--white)">
                <Heading size={900} marginTop="default">
                    {"📇 This page is under constructing"}
                </Heading>

            </Pane>
        </Pane>
    )
}