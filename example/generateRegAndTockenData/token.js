const ethUtil = require('ethereumjs-util');
const ethWallet = require('ethereumjs-wallet');

const priv = ethUtil.toBuffer(
	'0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19'
);
const pub = '04'.concat(ethUtil
	.bufferToHex(ethUtil.privateToPublic(priv))
	.slice(2));

const msg = ethUtil.toBuffer(
	[...(+new Date() + 48 * 60 * 60 * 1000).toString()].map((i) =>
		i.charCodeAt(0)
	)
);

const hash = ethUtil.keccak256(ethUtil.toBuffer(msg));

let sign = ethUtil.ecsign(hash, priv);
sign = ethUtil.bufferToHex(sign.r).slice(2) + ethUtil.bufferToHex(sign.s).slice(2);

console.log('token', `${ethUtil.bufferToHex(msg).slice(2)}.${sign}.${pub}`);