import * as React from "react";
import {Menu, Pane} from 'evergreen-ui'
import Link from 'next/link'

export const MenuBar = () => {
    return (
        <Pane width={185} backgroundColor="var(--white)" borderRight>
            <Menu>
                <Menu.Group>
                    <Link href={"/clients"}>
                        <Menu.Item icon="people">
                            {"Client"}
                        </Menu.Item>
                    </Link>
                    <Link href={"/payments"}>
                        <Menu.Item selected icon="exchange">
                            {"Payment"}
                        </Menu.Item>
                    </Link>
                    <Menu.Item icon="globe-network">{"Provider"}</Menu.Item>
                </Menu.Group>
                <Menu.Divider/>
                <Menu.Group>
                    <Menu.Item intent="danger">{"Logout"}</Menu.Item>
                </Menu.Group>
            </Menu>
        </Pane>
    )
}