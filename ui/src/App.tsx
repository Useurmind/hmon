import { Theme } from "@material-ui/core";
import * as styles from "@material-ui/styles";
import * as React from "react";

import { ThemedMasterBound } from "rfluxx-mui-theming";

import { ISiteMapStore } from "rfluxx-routing";
import { IPageManagementStore } from "rfluxx-routing";
import { CurrentPage } from "rfluxx-routing";

import { Master } from "./Master";


import { createTheme } from "./theming/Theme";


export interface IAppProps
{
    siteMapStore: ISiteMapStore;
    pageManagementStore: IPageManagementStore;
    
}

export interface IAppState
{
}

export class App extends React.Component<IAppProps, IAppState> {

    constructor(props: any)
    {
        super(props);

        this.state = {
        };
    }

    public render(): any
    {
        return (
            <CurrentPage pageManagementStore={this.props.pageManagementStore}
                        renderNoPage={() => <div className="container-fluid">
                            404: No page found on this url!</div>}
                        pageMasterTemplate={
                            <ThemedMasterBound storeRegistrationKey="IThemeStore"
                                                createTheme={createTheme}>
                            <Master />
                            </ThemedMasterBound>
                        } />
        );
    }
}
