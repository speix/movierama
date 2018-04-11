Vue.component('modal', {
    template: `
            <div class="modal is-active">
                <div class="modal-background"></div>
                <div class="modal-card">
                    <header class="modal-card-head">
                        <p class="modal-card-title">
                            <slot name="modal-title"></slot>
                        </p>
                        <button class="delete" aria-label="close" @click="$emit('close')"></button>
                    </header>
                    <section class="modal-card-body">
                        <slot name="modal-content"></slot>
                    </section>
                </div>
            </div>
            `
});

Vue.use(Toasted, {
    iconPack : 'fontawesome'
});