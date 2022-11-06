import Default from "./Default.svelte";
import Paragraph from "./Paragraph.svelte";
import Heading1 from "./Heading1.svelte";
import ChildDatabase from "./ChildDatabase.svelte";
import Divider from "./Divider.svelte";
import Callout from "./Callout.svelte";
import ChildPage from "./ChildPage.svelte";
import Heading2 from "./Heading2.svelte";
import Heading3 from "./Heading3.svelte";
import ColumnList from "./ColumnList.svelte";
import Column from "./Column.svelte";

export const components = {
    'heading_1': Heading1,
    'heading_2': Heading2,
    'heading_3': Heading3,
    "column_list": ColumnList,
    "column": Column,
    'paragraph': Paragraph,
    'child_database': ChildDatabase,
    'child_page': ChildPage,
    'divider': Divider,
    'callout': Callout,
}

export function loadComponent(name) {
    if (components[name]) {
        return components[name]
    }
    return Default
}