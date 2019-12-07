
/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
    public List<Integer> preorder(Node root) {

        List<Integer> list = new ArrayList<>();

        if(root == null) {
            return list;
        }

        Stack<Node> stack = new Stack<>();
        stack.add(root);
        while(!stack.isEmpty()) {
            root = stack.pop();
            int l = root.children.size();
            for(int j = l - 1; j >= 0; j--) {
                stack.add(root.children.get(j));
            }
            list.add(root.val);
        }

        return list;
    }
}