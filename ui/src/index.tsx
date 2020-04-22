
import * as React from "react";
import * as ReactDom from "react-dom";
import * as RfluxxRouting from "rfluxx-routing";
import { ISiteMapNode } from "rfluxx-routing";

import { App } from "./App";
import { siteMapNode as homeSiteMapNode } from "./pages/home/SiteMapNode";
import { GlobalContainerFactory } from "./GlobalContainerFactory";

const siteMap: ISiteMapNode = homeSiteMapNode;

const globalContainerFactory = new GlobalContainerFactory();

const globalStores = RfluxxRouting.init({
    siteMap,
    containerFactory: globalContainerFactory,
    targetNumberOpenPages: 5,
    rootPath: "ui"
});



document.addEventListener("DOMContentLoaded", event =>
{
    const root = document.getElementById("root");
    ReactDom.render(
        <App siteMapStore={globalStores.siteMapStore}
             pageManagementStore={globalStores.pageManagementStore}
             />,
        root);
});
