"""Spacemesh Proof-of-Space-Time (PoST) Verification."""
import hashlib

def verify_post_commitment(node_id: str, space_units: int, challenge: str) -> str:
    raw = f"{node_id}:{space_units}:{challenge}".encode()
    return hashlib.sha256(raw).hexdigest()
