import React, { useState } from 'react';
import {
  Container,
  Paper,
  Typography,
  Box,
  Divider,
  List,
  ListItem,
  ListItemText,
  CircularProgress
} from '@mui/material';
import { PayPalScriptProvider, PayPalButtons } from '@paypal/react-paypal-js';
import { useCart } from '../contexts/CartContext';
import { useNavigate } from 'react-router-dom';
import { useAuthenticatedFetch } from '../utils/api';

const PAYPAL_CLIENT_ID = "AdMk5sWX2p4KPXcwLUWiWaSa47mSqCepe-yHKPgFq7HzGfaVnLEgAWs_6YnEHhgqx86SQaQIXaXI50un";

const Checkout = () => {
  const { cart, total, clearCart } = useCart();
  const navigate = useNavigate();
  const fetchWithAuth = useAuthenticatedFetch();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [orderId, setOrderId] = useState(null);

  // Redirect if cart is empty
  if (cart.length === 0) {
    navigate('/cart');
    return null;
  }

  const createOrder = async () => {
    try {
      // Create order in our system first
      const response = await fetchWithAuth('/api/authenticated/v1/order', {
        method: 'POST',
        body: JSON.stringify({
          user_id: 1, // You might want to get this from your auth context
          total_amount: total,
          items: cart.map(item => ({
            menu_item_id: item.id,
            quantity: item.quantity,
            unit_price: item.price
          }))
        })
      });

      if (!response.success) {
        throw new Error('Failed to create order');
      }

      // Store the order ID for later use
      setOrderId(response.data.id);

      // Return the PayPal order creation
      return {
        purchase_units: [
          {
            amount: {
              value: total.toFixed(2),
              currency_code: "USD"
            },
            description: "Food Order"
          }
        ]
      };
    } catch (err) {
      setError('Failed to create order. Please try again.');
      throw err;
    }
  };

  const handlePaymentSuccess = async (data) => {
    setLoading(true);
    try {
      // Update order status to paid
      const response = await fetchWithAuth('/api/authenticated/v1/order/update_status', {
        method: 'POST',
        body: JSON.stringify({
          order_id: orderId,
          status: 'paid'
        })
      });

      if (!response.success) {
        throw new Error('Failed to update order status');
      }

      // Clear cart and redirect to success page
      clearCart();
      navigate('/order-success');
    } catch (err) {
      setError('Failed to process payment. Please try again.');
      console.error('Payment error:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <PayPalScriptProvider options={{ 
      "client-id": PAYPAL_CLIENT_ID,
      currency: "USD"
    }}>
      <Container maxWidth="md" sx={{ py: 4 }}>
        <Paper elevation={3} sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom>
            Order Summary
          </Typography>
          <List>
            {cart.map((item) => (
              <ListItem key={item.id}>
                <ListItemText
                  primary={item.name}
                  secondary={`Quantity: ${item.quantity}`}
                />
                <Typography>
                  ${(item.price * item.quantity).toFixed(2)}
                </Typography>
              </ListItem>
            ))}
          </List>
          <Divider sx={{ my: 2 }} />
          <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
            <Typography variant="h6">Total:</Typography>
            <Typography variant="h6">${total.toFixed(2)}</Typography>
          </Box>

          {error && (
            <Typography color="error" sx={{ mb: 2 }}>
              {error}
            </Typography>
          )}

          {loading ? (
            <Box sx={{ display: 'flex', justifyContent: 'center', my: 2 }}>
              <CircularProgress />
            </Box>
          ) : (
            <PayPalButtons
              createOrder={(data, actions) => {
                return createOrder().then(orderData => {
                  return actions.order.create(orderData);
                });
              }}
              onApprove={async (data, actions) => {
                const order = await actions.order.capture();
                await handlePaymentSuccess(order);
              }}
              onError={(err) => {
                setError('PayPal payment failed. Please try again.');
                console.error('PayPal Error:', err);
              }}
              style={{
                layout: "vertical",
                color: "gold",
                shape: "rect",
                label: "pay"
              }}
            />
          )}
        </Paper>
      </Container>
    </PayPalScriptProvider>
  );
};

export default Checkout; 