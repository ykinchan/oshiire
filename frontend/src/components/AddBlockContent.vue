<template>
	<div
		class="addBlock-content"
		v-if="isShowAddMenu == true"
		style="z-index: 2001;"
		@mousewheel.prevent
	>
		<!-- {{getterAddMenuContentLayerXY}} -->
		<div
			class="dropdown-menu"
			:style="{ top: getterAddMenuContentLayerXY.y, left:getterAddMenuContentLayerXY.x }"
			@mousewheel.stop
		>
			<div v-for="(item,index) in addBlockInfoArray" :key="index">
				<span class="block-type-tip" v-if="index == 0">BASIC BLOCKS</span>
				<span class="block-type-tip" v-if="index == 7">MEDIA</span>
				<div class="block-item" @click="addBlock(item.type)">
					<div class="block-item-img">
						<img :src="getImgUrl(item.type)" style="width: 100%;" />
					</div>
					<div class="block-item-intro">
						<h4>{{item.name}}</h4>
						<span>{{item.tip}}</span>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style lang="less">
.addBlock-content {
	position: fixed;
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	overflow: auto;
	margin: 0;
	.dropdown-menu {
		background: #ffffff;
		position: absolute;
		// margin-left: 58px;
		z-index: 2002;
		width: 320px;
		height: 360px;
		overflow-y: auto;
		border: 1px solid #dcdfe6;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 2px 12px 0px;
		border-radius: 4px;

		.block-type-tip {
			font-size: 12px;
			color: #909399;
			display: block;
			padding: 10px 20px 5px 20px;
		}
		.block-item {
			padding: 5px 20px 7px 20px;
			display: flex;
			align-items: center;
			background: #ffffff;
			.block-item-img {
				width: 45px;
				height: 45px;
				background: #ffffff;
				img {
					border-radius: 4px;
					border: 1px solid #dddddd;
				}
			}
			.block-item-intro {
				h4 {
					margin: 0;
					font-size: 14px;
				}
				span {
					font-size: 12px;
					color: #909399;
				}
				margin-left: 10px;
			}
		}
		.block-item:hover {
			background: #eeeeee;
			cursor: pointer;
		}
	}
}
</style>


<script>
// @ is an alias to /src

export default {
	name: "addBlock-content",
	data() {
		return {
			isShowMenu: this.isShowAddMenu,
			addBlockInfoArray: [
				{
					name: "Text",
					tip: "Just start writing with plain text.",
					type: "text"
				},
				{
					name: "To-do list",
					tip: "Track tasks with a to-do list.",
					type: "todo"
				},
				{
					name: "Heading 1",
					tip: "Big section heading.",
					type: "heading1"
				},
				{
					name: "Heading 2",
					tip: "Medium section heading.",
					type: "heading2"
				},
				{
					name: "Heading 3",
					tip: "Small section heading.",
					type: "heading3"
				},
				{
					name: "Bulleted list",
					tip: "Create a simple bulleted list",
					type: "BulletedList"
				},
				{
					name: "Quote",
					tip: "Capture a quote",
					type: "hint"
				},
				{
					name: "Quote",
					tip: "Capture a quote",
					type: "hint"
				},
				{
					name: "Image",
					tip: "select imgae",
					type: "image"
				}
			]
		};
	},
	watch: {
		isShowAddMenu: function(value) {
			if (value == true) {
				document.addEventListener("click", e => {
					// console.log(event.target.getAttribute("class"))
					if (
						event.target.getAttribute("class") !=
						"iconfont iconplus"
					) {
						if (
							event.target.getAttribute("class") !=
							"dropdown-menu"
						) {
							this.$store.commit("mutationIsShowAddMenu", false);
						}
					}
				});
			}
		}
	},
	computed: {
		isShowAddMenu() {
			return this.$store.state.isShowAddMenu;
		},
		currentBlockIndex() {
			return this.$store.state.currentBlockIndex;
		},
		getterAddMenuContentLayerXY() {
			return this.$store.getters.getterAddMenuContentLayerXY;
		},
		currentPageBlocks() {
			return this.$store.state.currentPageBlocks;
		}
	},
	methods: {
		addBlock(a) {
			let addBlockInfo = {
				index: this.currentBlockIndex,
				blockItem: {}
			};
			if (a == "text") {
				addBlockInfo.blockItem = {
					type: "text",
					data: {
						text: ""
					}
				};
			}
			if (a == "todo") {
				addBlockInfo.blockItem = {
					type: "todo",
					data: {
						isChecked: false,
						text: ""
					}
				};
			}
			if (a == "heading1") {
				addBlockInfo.blockItem = {
					type: "heading1",
					data: {
						text: ""
					}
				};
			}
			if (a == "heading2") {
				addBlockInfo.blockItem = {
					type: "heading2",
					data: {
						text: ""
					}
				};
			}
			if (a == "heading3") {
				addBlockInfo.blockItem = {
					type: "heading3",
					data: {
						text: ""
					}
				};
			}
			if (a == "BulletedList") {
				addBlockInfo.blockItem = {
					type: "BulletedList",
					data: {
						text: ""
					}
				};
			}
			if (a == "hint") {
				addBlockInfo.blockItem = {
					type: "hint",
					data: {
						text: ""
					}
				};
			}
			if (a == "image") {
				addBlockInfo.blockItem = {
					type: "image",
					data: {
						text: ""
					}
				};
			}

			this.$store.commit("mutationAddCurrentPageBlocks", addBlockInfo);

			// ??????????????????????????????????????????text???????????????????????????????????????????????????????????????
			if (
				this.currentPageBlocks[this.currentBlockIndex].type == "text" &&
				this.currentPageBlocks[this.currentBlockIndex].data.text == ""
			) {
				// ???????????????????????????????????????????????????
				this.currentPageBlocks.splice(this.currentBlockIndex, 1);
				setTimeout(() => {
					let dom = document.getElementsByTagName("textarea");
					let currInput = dom[this.currentBlockIndex];
					currInput.focus();
				}, 300);
			} else {
				// ???????????????????????????????????????????????????????????????
				setTimeout(() => {
					let dom = document.getElementsByTagName("textarea");
					let nextInput = dom[this.currentBlockIndex + 1];
					nextInput.focus();
				}, 300);
			}
		},
		getImgUrl(type) {
			return require("@/assets/" + type + ".png");
		}
	}
};
</script>


