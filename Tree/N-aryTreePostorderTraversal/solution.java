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
    public List<Integer> postorder(Node root) {
        LinkedList<Integer> list = new LinkedList<>();
        if(root == null) {
            return list;
        }

        Stack<Node> stack = new Stack<>();
        stack.add(root);
        while(!stack.isEmpty()) {
            Node node = stack.pop();
            list.addFirst(node.val);
            for(Node item: node.children) {
                if(item != null) {
                    stack.add(item);
                }
            }
        }

        return list;
    }
}