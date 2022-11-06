import Title from "./Title.svelte";
import Number from "./Number.svelte";
import Select from "./Select.svelte";
import MultiSelect from "./MultiSelect.svelte";
import Text from "./Text.svelte";
import Default from "./Default.svelte";
import Textarea from "./Textarea.svelte";
import Email from "./Email.svelte";
import Phone from "./Phone.svelte";
import Date from "./Date.svelte";
import Checkbox from "./Checkbox.svelte";

export const components = {
    'title': Title,
    'number': Number,
    'select': Select,
    'multi_select': MultiSelect,
    'rich_text': Textarea,
    'url': Text,
    'email': Email,
    'phone_number': Phone,
    'date': Date,
    'checkbox':Checkbox,
}

export function resolveField(name) {
    if (components[name]) {
        return components[name]
    }
    return Default
}