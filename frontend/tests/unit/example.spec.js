import { expect } from 'chai';
import { createLocalVue, shallowMount } from '@vue/test-utils';
import Vue from 'vue';
import Vuex from 'vuex';
import flushPromises from 'flush-promises';
import HelloWorld from '@/components/HelloWorld.vue';

const localVue = createLocalVue();
localVue.use(Vuex);

describe('HelloWorld.vue', () => {
  const expected = ['a', 'b', 'c'];

  let actions;
  let store;
  beforeEach(() => {
    actions = {
      getComments: () => ({
        data: expected,
      }),
    };
    store = new Vuex.Store({
      state: {},
      actions,
    });
  });

  it('값을 잘 읽어옴', async () => {
    const wrapper = shallowMount(HelloWorld, {
      localVue,
      store,
    });

    // NOTE - created 이후 화면 갱신 사이클 대기
    await Vue.nextTick();
    await flushPromises();
    expect(JSON.parse((wrapper.get('label')
      .text())))
      .to
      .deep
      .equal(expected);
  });
});
