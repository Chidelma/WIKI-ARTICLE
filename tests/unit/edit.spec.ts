import { mount } from "@vue/test-utils";
import Edit from "@/views/Edit.vue";

describe("Edit", () => {

    const wrapper = mount(Edit, {
        props: {
            name: "rest_api"
        }
    });

    test('article name for edit', () => {

        const header = wrapper.find("#article");

        expect(header.text()).toBe("REST_API");
    });

    test('article description length', () => {

        const desc = wrapper.findAll('.desc');

        expect(desc.length).toBe(2);
    });

    test('article input description', () => {

        const desc = wrapper.find('#input-desc');

        setTimeout(() => {
            expect(desc.text()).toBe("REpresentational State Transfer Application Programming Interface");
        }, 1000);
    });

    test('article preview description', () => {

        const desc = wrapper.find('#preview');

        setTimeout(() => {
            expect(desc.text()).toBe("REpresentational State Transfer Application Programming Interface");
        }, 1000);
    });
});