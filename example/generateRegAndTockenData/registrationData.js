const ethUtil   = require('ethereumjs-util');
const ethWallet = require('ethereumjs-wallet');

const priv = ethUtil.toBuffer('0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19');
const pub  = ethUtil
	.bufferToHex(new Uint8Array(ethUtil.privateToPublic(priv)))
	.slice(2);

const msg = ethUtil.toBuffer(
	[
		...JSON.stringify([
			{
				network: 'Handshake',
				coin: 'HSN',
				address: 'hs1qyvu2uyh8kqgxts0dacy0m47at005nrxn2zuqnl',
			},
		]),
	].map((i) => i.charCodeAt(0))
);

hash = ethUtil.keccak256(ethUtil.toBuffer(msg));

sign = ethUtil.ecsign(hash, priv);
sign = ethUtil.bufferToHex(sign.r).slice(2) + ethUtil.bufferToHex(sign.s).slice(2);

console.log({
	addresses: ethUtil.bufferToHex(msg).slice(2),
	sign,
	pubKey: pub,
});